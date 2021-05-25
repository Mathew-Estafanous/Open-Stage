package handler

import (
	"encoding/json"
	"fmt"
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

func (c *CreateAccount) UnmarshalJSON(data []byte) error {
	type create2 CreateAccount
	if err := json.Unmarshal(data, (*create2)(c)); err != nil {
		return err
	}

	if c.Name == "" {
		return fmt.Errorf("%w: missing account 'name' field", domain.BadInput)
	}
	if c.Email == "" {
		return fmt.Errorf("%w: missing account 'email' field", domain.BadInput)
	}
	if c.Username == "" {
		return fmt.Errorf("%w: missing account 'username' field", domain.BadInput)
	}
	if c.Password == "" {
		return fmt.Errorf("%w: missing account 'password' field", domain.BadInput)
	}
	return nil
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

	resp := AccountResp{
		Id: acc.Id,
		Name: acc.Name,
		Username: acc.Username,
		Email: acc.Email,
	}
	a.respond(w, http.StatusCreated, resp)
}
