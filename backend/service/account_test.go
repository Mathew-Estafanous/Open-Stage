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

	err := service.Delete(1, 1)
	assert.NoError(t, err)

	err = service.Delete(1, 0)
	assert.ErrorIs(t, err, domain.Forbidden)
}

func TestAccountService_FindByUsername(t *testing.T) {
	aStore := new(mock.AccountStore)
	service := NewAccountService(aStore)

	acc := domain.Account{
		Id:       10,
		Username: "TheUsername",
		Name:     "Jackson",
		Password: "A-PASSWORD",
		Email:    "jackson@fake.com",
	}
	aStore.On("GetByUsername", acc.Username).Return(acc, nil)

	account, err := service.FindByUsername(acc.Username, acc.Id)
	assert.NoError(t, err)
	assert.EqualValues(t, acc, account)

	aStore.On("GetByUsername", "invalidUsername").Return(domain.Account{}, domain.NotFound)

	_, err = service.FindByUsername("invalidUsername", acc.Id)
	assert.ErrorIs(t, err, domain.NotFound)

	_, err = service.FindByUsername(acc.Username, 1)
	assert.ErrorIs(t, err, domain.Forbidden)
}
