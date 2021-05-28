package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// AccountResp represents a user's account information like name,
// username, etc.
//
// swagger:model accountResponse
type AccountResp struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateAccount are the fields that are used to signup a new user account.
//
// swagger:model createAccount
type CreateAccount struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UnmarshalJSON will enforce all the required fields in the CreateAccount
// struct and returns an error if a field is missing.
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
	r.HandleFunc("/accounts/signup", a.createAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", a.deleteAccount).Methods("DELETE")
}

// swagger:route POST /accounts Accounts createAccount
//
// Signup a new user account.
//
// Will create a new account while also ensuring the validity of the provided email.
//
// Responses:
//   201: accountResponse
//   400: errorResponse
//   409: errorResponse
//   500: errorResponse
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

// swagger:route DELETE /accounts/{accountId} Accounts accountId
//
// Delete an account by ID
//
// Will delete the user account with the correlating account 'id.'
//
// Responses:
//   200: description: OK - Question has been properly deleted.
//   400: errorResponse
//   500: errorResponse
func (a accountHandler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	pathId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(pathId)
	if err != nil {
		a.error(w, fmt.Errorf("%w: given id is not a valid int", domain.BadInput))
		return
	}

	err = a.as.Delete(id)
	if err != nil {
		a.error(w, err)
		return
	}
}
