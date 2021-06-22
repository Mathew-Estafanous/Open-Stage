package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestAuthService_Authenticate(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	aStore := new(mock.AccountStore)
	cache := new(mock.AuthCache)
	auth := NewAuthService(aStore, cache)

	respAcc := domain.Account{
		Username: "someUsername",
		// This is a hashed password 'helloWorld' using bcrypt.
		Password: "$2y$12$UvUX39hRbeEGVCgEZmX3NO/5No10LEFe7ZsARJ5iK/55oSOUs7Bha",
	}
	aStore.On("GetByUsername", respAcc.Username).Return(respAcc, nil)

	username := "someUsername"
	password := "helloWorld"
	authToken, err := auth.Authenticate(username, password)
	assert.NoError(t, err)

	validateAuthTkn(authToken, t)

	username = "InvalidUsername"
	aStore.On("GetByUsername", username).Return(domain.Account{}, domain.NotFound)
	_, err = auth.Authenticate(username, password)
	assert.Error(t, err)
}

func TestAuthService_Refresh(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	aStore := new(mock.AccountStore)
	cache := new(mock.AuthCache)
	auth := NewAuthService(aStore, cache)

	exp := time.Now().Add(time.Hour * 168).Unix()
	refreshClaim := AccountClaims{
		"USERNAME",
		jwt.StandardClaims{
			ExpiresAt: exp,
			Audience:  "refresh",
			Subject:   "2",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshTkn, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)
	cache.On("Contains", refreshTkn).Return(false, nil)

	authTkn, err := auth.Refresh(refreshTkn)
	assert.NoError(t, err)
	validateAuthTkn(authTkn, t)

	exp = time.Date(2000, 1, 1, 1, 0, 0, 0, time.UTC).Unix()
	refreshClaim.ExpiresAt = exp
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	expiredRefreshTkn, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	assert.NoError(t, err)
	cache.On("Contains", expiredRefreshTkn).Return(false, nil)

	_, err = auth.Refresh(expiredRefreshTkn)
	assert.ErrorIs(t, err, domain.Unauthorized)
}

func TestAuthService_Invalidate(t *testing.T) {
	aStore := new(mock.AccountStore)
	cache := new(mock.AuthCache)
	auth := NewAuthService(aStore, cache)

	authTkn := domain.AuthToken{
		AccessToken:  "SomeAccessJWT",
		RefreshToken: "TheRefreshJWT",
	}
	cache.On("Store", authTkn.AccessToken).Return(nil)
	cache.On("Store", authTkn.RefreshToken).Return(nil)

	err := auth.Invalidate(authTkn)
	assert.NoError(t, err)
}

func validateAuthTkn(authToken domain.AuthToken, t *testing.T) {
	accessTkn, err := jwt.Parse(authToken.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	assert.NoError(t, err)

	refreshTkn, err := jwt.Parse(authToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	assert.NoError(t, err)

	assert.EqualValues(t, true, accessTkn.Valid)
	assert.EqualValues(t, true, refreshTkn.Valid)
}
