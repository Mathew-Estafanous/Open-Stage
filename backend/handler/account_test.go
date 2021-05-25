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

	account := domain.Account{
		Name: "Mathew",
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email: "mathew@gmail.com",
	}
	as.On("Create", &account).Return(nil)

	createAcc := CreateAccount{
		Name: "Mathew",
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email: "mathew@gmail.com",
	}

	j, err := json.Marshal(createAcc)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/accounts", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	NewAccountHandler(as).Route(r)
	r.ServeHTTP(w, req)

	resp := AccountResp{
		Id: 0,
		Name: "Mathew",
		Username: "MatMat",
		Email: "mathew@gmail.com",
	}
	j, err = json.Marshal(resp)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())


	missingField := CreateAccount{
		// Missing the 'name' field
		Username: "MatMat",
		Password: "ThisIsAPassword",
		Email: "mathew@gmail.com",
	}
	j, err = json.Marshal(missingField)
	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/accounts", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}
