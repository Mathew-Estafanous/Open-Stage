package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience: "access",
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})
	accessTkn, err := tkn.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)

	var passedAuth bool
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		passedAuth = true
	})
	auth := Auth(h)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer " + accessTkn)

	auth.ServeHTTP(w, req)
	assert.EqualValues(t, true, passedAuth)


	tkn = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience: "refresh",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	refreshTkn, err := tkn.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)

	passedAuth = false
	req.Header.Set("Authorization", "Bearer " + refreshTkn)

	auth.ServeHTTP(w, req)
	assert.EqualValues(t, false, passedAuth)
}
