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

// Refresh contains the provided refresh token that is used to retrieve a new
// access token.
//
// swagger:model refresh
type Refresh struct {
	Tkn string `json:"refresh_token"`
}

type accountHandler struct {
	baseHandler
	as   domain.AccountService
	auth domain.AuthService
}

func NewAccountHandler(aService domain.AccountService, authService domain.AuthService) *accountHandler {
	return &accountHandler{
		as:   aService,
		auth: authService,
	}
}

func (a accountHandler) Route(r, secured *mux.Router) {
	r.HandleFunc("/accounts/signup", a.createAccount).Methods("POST")
	r.HandleFunc("/accounts/login", a.login).Methods("POST")
	r.HandleFunc("/accounts/logout", a.logout).Methods("POST")
	r.HandleFunc("/accounts/refresh", a.refresh).Methods("POST")

	secured.HandleFunc("/accounts/{id}", a.deleteAccount).Methods("DELETE")
	secured.HandleFunc("/accounts/{username}", a.findWithUsername).Methods("GET")
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

	a.respond(w, http.StatusCreated, accountToResp(acc))
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

// swagger:route GET /accounts/{username} Accounts accountUsername
//
// Get account information by username.
//
// Retrieve relevant account information by providing the username.
//
// Responses:
//  200: accountResponse
//  403: errorResponse
//  404: errorResponse
//  500: errorResponse
func (a accountHandler) findWithUsername(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	accId, err := strconv.Atoi(r.Header.Get("Account"))
	if err != nil {
		a.error(w, err)
		return
	}

	acc, err := a.as.FindByUsername(username, accId)
	if err != nil {
		a.error(w, err)
		return
	}

	a.respond(w, http.StatusOK, accountToResp(acc))
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

// swagger:route POST /accounts/refresh Accounts refresh
//
// Refresh current access token.
//
// Use your provided refresh token to get another access token. Remember that refresh
// tokens also expire and must be used within their expiration time. If the refresh
// token is expired, you must authenticate again.
//
// Responses:
//  200: authToken
//  401: errorResponse
//  500: errorResponse
func (a accountHandler) refresh(w http.ResponseWriter, r *http.Request) {
	var body Refresh
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		a.error(w, err)
		return
	}

	token, err := a.auth.Refresh(body.Tkn)
	if err != nil {
		a.error(w, err)
		return
	}

	a.respond(w, http.StatusOK, token)
}

// swagger:route POST /accounts/logout Accounts authTokens
//
// Logout from current account credentials.
//
// Takes both access and refresh token credentials and adds them to a blacklist. Those tokens
// will remain blacklisted until their expiration date. Meaning that they can no longer be used
// for secured endpoints.
//
// Responses:
//  200: description: OK - Account credentials have been blacklisted.
//  401: errorResponse
//  500: errorResponse
func (a accountHandler) logout(w http.ResponseWriter, r *http.Request) {
	var body domain.AuthToken
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		a.error(w, err)
		return
	}

	err = a.auth.Invalidate(body)
	if err != nil {
		a.error(w, err)
		return
	}
}

func accountToResp(acc domain.Account) AccountResp {
	return AccountResp{
		Id:       acc.Id,
		Name:     acc.Name,
		Username: acc.Username,
		Email:    acc.Email,
	}
}
