package mock

import "github.com/stretchr/testify/mock"

type AuthService struct {
	mock.Mock
}

func (a *AuthService) OwnsRoom(code string, accId int) (bool, error) {
	ret := a.Called(code, accId)
	return ret.Get(0).(bool), ret.Error(1)
}

