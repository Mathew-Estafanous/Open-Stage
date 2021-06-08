package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccountHandler_createAccount(t *testing.T) {
	as := new(mock.AccountService)
	auth := new(mock.AuthService)

	account := domain.Account{
		Name:     "Mathew",
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email:    "mathew@gmail.com",
	}
	as.On("Create", &account).Return(nil)

	createAcc := CreateAccount{
		Name:     "Mathew",
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email:    "mathew@gmail.com",
	}

	j, err := json.Marshal(createAcc)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/accounts/signup", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	secured := mock.SecureRouter(r, 1)
	NewAccountHandler(as, auth).Route(r, secured)
	r.ServeHTTP(w, req)

	resp := AccountResp{
		Id:       0,
		Name:     "Mathew",
		Username: "MatMat",
		Email:    "mathew@gmail.com",
	}
	j, err = json.Marshal(resp)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	missingField := CreateAccount{
		// Missing the 'name' field
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email:    "mathew@gmail.com",
	}
	j, err = json.Marshal(missingField)
	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/accounts/signup", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestAccountHandler_deleteAccount(t *testing.T) {
	as := new(mock.AccountService)
	auth := new(mock.AuthService)

	as.On("Delete", 5, 5).Return(nil)

	req, err := http.NewRequest("DELETE", "/accounts/5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	secured := mock.SecureRouter(r, 5)
	NewAccountHandler(as, auth).Route(r, secured)

	r.ServeHTTP(w, req)
	assert.EqualValues(t, http.StatusOK, w.Code)

	req, err = http.NewRequest("DELETE", "/accounts/sfaf", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestAccountHandler_login(t *testing.T) {
	as := new(mock.AccountService)
	auth := new(mock.AuthService)


	login := Login{
		Username: "Mathew",
		Password: "SecretPassword",
	}

	tks := domain.AuthToken{
		AccessToken: "AN-ACCESS-TOKEN",
		RefreshToken: "A-REFRESH-TOKEN",
	}

	auth.On("Authenticate", login.Username, login.Password).Return(tks, nil)

	j, err := json.Marshal(login)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/accounts/login", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	secured := mock.SecureRouter(r, 5)
	NewAccountHandler(as, auth).Route(r, secured)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	j, err = json.Marshal(tks)
	assert.NoError(t, err)

	assert.JSONEq(t, string(j), w.Body.String())
}
