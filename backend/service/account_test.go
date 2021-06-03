package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
)

func TestAccountService_Create(t *testing.T) {
	aStore := new(mock.AccountStore)
	service := NewAccountService(aStore)

	account := domain.Account{
		Name:     "Hello",
		Username: "Mathew",
		Password: "MatMat",
		Email:    "mathew@gmail.com",
	}

	aStore.On("Create", mock2.MatchedBy(func(a *domain.Account) bool {
		err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte("MatMat"))
		return err == nil
	})).Return(nil)

	err := service.Create(&account)
	assert.NoError(t, err)

	invAccount := domain.Account{
		Name:     "Hello",
		Username: "Mathew",
		Password: "MatMat",
		Email:    "mathew@not-a-valid-address.com",
	}

	err = service.Create(&invAccount)
	assert.ErrorIs(t, err, domain.BadInput)
}

func TestAccountService_Delete(t *testing.T) {
	aStore := new(mock.AccountStore)
	service := NewAccountService(aStore)

	aStore.On("Delete", 1).Return(nil)

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestAccountService_Authenticate(t *testing.T) {
	err := os.Setenv("SECRET_KEY", "SECRET")
	assert.NoError(t, err)

	aStore := new(mock.AccountStore)
	service := NewAccountService(aStore)

	respAcc := domain.Account{
		Username: "someUsername",
		// This is a hashed password 'helloWorld' using bcrypt.
		Password: "$2y$12$UvUX39hRbeEGVCgEZmX3NO/5No10LEFe7ZsARJ5iK/55oSOUs7Bha",
	}
	aStore.On("GetByUsername", respAcc.Username).Return(respAcc, nil)

	acc := domain.Account{
		Username: "someUsername",
		Password: "helloWorld",
	}
	authToken, err := service.Authenticate(acc)
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

	acc.Username = "InvalidUsername"
	aStore.On("GetByUsername", acc.Username).Return(domain.Account{}, domain.NotFound)
	_, err = service.Authenticate(acc)
	assert.Error(t, err)
}
