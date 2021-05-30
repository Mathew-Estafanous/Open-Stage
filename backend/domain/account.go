package domain

type Account struct {
	Id       int
	Name     string
	Username string
	Password string
	Email    string
}

// AuthToken contains both the access and refresh tokens after
// a user has successfully authenticated.
//
// swagger:model authToken
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccountStore interface {
	Create(acc *Account) error
	GetByUsername(username string) (Account, error)
	Delete(id int) error
}

type AccountService interface {
	Create(acc *Account) error
	Authenticate(acc Account) (AuthToken, error)
	Delete(id int) error
}
