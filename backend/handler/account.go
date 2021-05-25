package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountResp struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateAccount struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type accountHandler struct {
	baseHandler
	as domain.AccountService
}

func NewAccountHandler(aService domain.AccountService) *accountHandler {
	return &accountHandler{as: aService}
}

func (a accountHandler) Route(r *mux.Router) {
	r.HandleFunc("/accounts", a.createAccount).Methods("POST")
}

func (a accountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var body *CreateAccount
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		a.error(w, err)
		return
	}

	acc := domain.Account{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}

	err = a.as.Create(&acc)
	if err != nil {
		a.error(w, err)
		return
	}

	resp := mapAccToResp(acc)
	a.respond(w, http.StatusCreated, resp)
}

func mapAccToResp(acc domain.Account) AccountResp {
	return AccountResp{
		Id: acc.Id,
		Name: acc.Name,
		Username: acc.Username,
		Email: acc.Email,
	}
}
