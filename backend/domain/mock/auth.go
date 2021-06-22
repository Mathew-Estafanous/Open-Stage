package mock

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"strconv"
)

type AuthService struct {
	mock.Mock
}

func (a *AuthService) Authenticate(username, password string) (domain.AuthToken, error) {
	ret := a.Called(username, password)
	return ret.Get(0).(domain.AuthToken), ret.Error(1)
}

func (a *AuthService) Refresh(refreshTkn string) (domain.AuthToken, error) {
	ret := a.Called(refreshTkn)
	return ret.Get(0).(domain.AuthToken), ret.Error(1)
}

func (a *AuthService) Invalidate(token domain.AuthToken) error {
	ret := a.Called(token)
	return ret.Error(0)
}

func SecureRouter(r *mux.Router, fakeId int) *mux.Router {
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Account", strconv.Itoa(fakeId))
			h.ServeHTTP(w, r)
		})
	})
	return secured
}

type AuthCache struct {
	mock.Mock
}

func (a *AuthCache) Contains(tkn string) (bool, error) {
	ret := a.Called(tkn)
	return ret.Get(0).(bool), ret.Error(1)
}

func (a *AuthCache) Store(tkn string) error {
	ret := a.Called(tkn)
	return ret.Error(0)
}
