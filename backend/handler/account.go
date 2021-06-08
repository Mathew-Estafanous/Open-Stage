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
	// The account id.
	//
	// example: 1234
	Id int `json:"id"`

	// Name of the account holder.
	//
	// example: Mathew
	Name string `json:"name"`

	// The account username.
	//
	// example: MatMat2
	Username string `json:"username"`

	// Email associated with the account.
	//
	// example: mathew@fake.com
	Email string `json:"email"`
}

// CreateAccount are the fields that are used to signup a new user account.
//
// swagger:model createAccount
type CreateAccount struct {
	// required: true
	// example: Mathew
	Name string `json:"name"`

	// required: true
	// example: MatMat
	Username string `json:"username"`

	// required: true
	// example: aSecretPassword
	Password string `json:"password"`

	// required: true
	// example: mathew@fake.com
	Email string `json:"email"`
}

// Login are the fields that are required to successfully log into an account.
//
// swagger:model loginAccount
type Login struct {
	// Username of the account wanted.
	//
	// required: true
	// example: MatMat
	Username string `json:"username"`

	// Password of the account.
	//
	// required: true
	// example: aSecretPassword
	Password string `json:"password"`
}

func (l *Login) UnmarshalJSON(data []byte) error {
	type login2 Login
	if err := json.Unmarshal(data, (*login2)(l)); err != nil {
		return err
	}

	if l.Username == "" {
		return fmt.Errorf("%w: missing account 'username' field", domain.BadInput)
	}
	if l.Password == "" {
		return fmt.Errorf("%w: missing account 'password' field", domain.BadInput)
	}
	return nil
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
	auth domain.AuthService
}

func NewAccountHandler(aService domain.AccountService, authService domain.AuthService) *accountHandler {
	return &accountHandler{
		as: aService,
		auth: authService,
	}
}

func (a accountHandler) Route(r, secured *mux.Router) {
	r.HandleFunc("/accounts/signup", a.createAccount).Methods("POST")
	r.HandleFunc("/accounts/login", a.login).Methods("POST")

	secured.HandleFunc("/accounts/{id}", a.deleteAccount).Methods("DELETE")
}

// swagger:route POST /accounts/signup Accounts createAccount
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
		Id:       acc.Id,
		Name:     acc.Name,
		Username: acc.Username,
		Email:    acc.Email,
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
//   403: errorResponse
//   500: errorResponse
func (a accountHandler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	pathId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(pathId)
	if err != nil {
		a.error(w, fmt.Errorf("%w: given id is not a valid int", domain.BadInput))
		return
	}

	accId, err := strconv.Atoi(r.Header.Get("Account"))
	if err != nil {
		a.error(w, err)
		return
	}
	err = a.as.Delete(id, accId)
	if err != nil {
		a.error(w, err)
		return
	}
}

// swagger:route POST /accounts/login Accounts loginAccount
//
// Login and authenticate account.
//
// Uses the provided credentials to authenticate and will return JWT tokens
// if authentication is successful.
//
// Responses:
//   200: authToken
//   401: errorResponse
//   500: errorResponse
func (a accountHandler) login(w http.ResponseWriter, r *http.Request) {
	var body Login
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		a.error(w, err)
		return
	}

	token, err := a.auth.Authenticate(body.Username, body.Password)
	if err != nil {
		a.error(w, err)
		return
	}

	a.respond(w, http.StatusOK, token)
}
