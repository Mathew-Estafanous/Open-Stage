package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAuthService_Authenticate(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	aStore := new(mock.AccountStore)
	auth := NewAuthService(aStore)

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

	accessTkn, err := jwt.Parse(authToken.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	assert.NoError(t, err)

	refreshTkn, err := jwt.Parse(authToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	assert.NoError(t, err)

	assert.EqualValues(t, true, accessTkn.Valid)
	assert.EqualValues(t, true, refreshTkn.Valid)

	username = "InvalidUsername"
	aStore.On("GetByUsername", username).Return(domain.Account{}, domain.NotFound)
	_, err = auth.Authenticate(username, password)
	assert.Error(t, err)
}