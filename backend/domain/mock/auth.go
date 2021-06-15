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

func SecureRouter(r *mux.Router, fakeId int) *mux.Router {
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("AccountPage", strconv.Itoa(fakeId))
			h.ServeHTTP(w, r)
		})
	})
	return secured
}

func (a *AuthService) Refresh(refreshTkn string) (domain.AuthToken, error) {
	ret := a.Called(refreshTkn)
	return ret.Get(0).(domain.AuthToken), ret.Error(1)
}
