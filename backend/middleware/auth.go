package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/handle_err"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

func Auth(cache domain.AuthCache) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			if len(auth) != 2 {
				writeError(w, handle_err.ToHttp(NotFormatted))
				return
			}

			blacklisted, err := cache.Contains(auth[1])
			if err != nil {
				writeError(w, handle_err.ToHttp(BlacklistErr))
				return
			}
			if blacklisted {
				writeError(w, handle_err.ToHttp(Blacklisted))
				return
			}

			tk, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_KEY")), nil
			})
			if err != nil {
				writeError(w, handle_err.ToHttp(InvalidToken))
				return
			}

			c := tk.Claims.(jwt.MapClaims)
			if c["aud"] != "access" {
				writeError(w, handle_err.ToHttp(InvalidToken))
				return
			}

			accId, ok := c["sub"].(string)
			if !ok {
				writeError(w, handle_err.ToHttp(Unidentified))
				return
			}

			r.Header.Set("Account", accId)
			next.ServeHTTP(w, r)
		})
	}
}

var (
	InvalidToken = fmt.Errorf("%w: Invalid access token", domain.Unauthorized)
	Blacklisted  = fmt.Errorf("%w: Provided token has been blacklisted", domain.Unauthorized)
	Unidentified = fmt.Errorf("%w: Unable to identify the account holder", domain.Unauthorized)
	NotFormatted = fmt.Errorf("%w: Authorization header not formatted correctly", domain.Unauthorized)
	BlacklistErr = fmt.Errorf("%w: Unable to get token blacklist", domain.Internal)
)

func writeError(w http.ResponseWriter, respErr handle_err.ResponseError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respErr.Sts)
	err := json.NewEncoder(w).Encode(respErr)
	if err != nil {
		log.Print(err)
	}
}
