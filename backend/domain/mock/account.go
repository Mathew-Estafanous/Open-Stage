package mock

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/mock"
)

type AccountService struct {
	mock.Mock
}

func (a *AccountService) Create(acc *domain.Account) error {
	ret := a.Called(acc)
	return ret.Error(0)
}

func (a *AccountService) Delete(id, accId int) error {
	ret := a.Called(id, accId)
	return ret.Error(0)
}

type AccountStore struct {
	mock.Mock
}

func (a *AccountStore) Create(acc *domain.Account) error {
	ret := a.Called(acc)
	return ret.Error(0)
}

func (a *AccountStore) GetByUsername(username string) (domain.Account, error) {
	ret := a.Called(username)
	return ret.Get(0).(domain.Account), ret.Error(1)
}

func (a *AccountStore) Delete(id int) error {
	ret := a.Called(id)
	return ret.Error(0)
}
