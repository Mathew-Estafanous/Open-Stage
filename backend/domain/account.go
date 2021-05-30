package domain

type Account struct {
	Id       int
	Name     string
	Username string
	Password string
	Email    string
}

type Token struct {
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
	Authenticate(acc Account) (Token, error)
	Delete(id int) error
}
