package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
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
