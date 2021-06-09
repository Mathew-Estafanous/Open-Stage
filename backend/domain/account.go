package domain

type Account struct {
	Id       int
	Name     string
	Username string
	Password string
	Email    string
}

type AccountStore interface {
	Create(acc *Account) error
	GetByUsername(username string) (Account, error)
	Delete(id int) error
}

type AccountService interface {
	Create(acc *Account) error
	Delete(id, accId int) error
}
