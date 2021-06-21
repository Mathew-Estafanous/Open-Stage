package middleware

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Auth(cache domain.AuthCache) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			if len(auth) != 2 {
				writeError(w, "Authorization header not formatted correctly.")
				return
			}

			blacklisted, _ := cache.Contains(auth[1])
			if blacklisted {
				writeError(w, "Provided token has been blacklisted.")
				return
			}

			tk, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_KEY")), nil
			})
			if err != nil {
				writeError(w, "Invalid access token.")
				return
			}

			c := tk.Claims.(jwt.MapClaims)
			if c["aud"] != "access" {
				writeError(w, "Invalid access token")
				return
			}

			accId, ok := c["sub"].(string)
			if !ok {
				writeError(w, "Unable to identify the account holder")
				return
			}

			r.Header.Set("Account", accId)
			next.ServeHTTP(w, r)
		})
	}
}

type response struct {
	Msg       string    `json:"message"`
	Sts       int       `json:"status"`
	TimeStamp time.Time `json:"timestamp"`
}

func writeError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respErr := response{
		Msg:       msg,
		Sts:       http.StatusUnauthorized,
		TimeStamp: time.Now(),
	}
	w.WriteHeader(respErr.Sts)
	err := json.NewEncoder(w).Encode(respErr)
	if err != nil {
		log.Print(err)
	}
}
