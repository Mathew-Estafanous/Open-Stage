package mock

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"strconv"
)

type AuthService struct {
	mock.Mock
}

func (a *AuthService) OwnsRoom(code string, accId int) (bool, error) {
	ret := a.Called(code, accId)
	return ret.Get(0).(bool), ret.Error(1)
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

