package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"golang.org/x/crypto/bcrypt"
	"net"
	"regexp"
	"strings"
)

type accountService struct {
	store domain.AccountStore
}

func NewAccountService(aStore domain.AccountStore) domain.AccountService {
	return &accountService{
		store: aStore,
	}
}

func (a *accountService) Create(acc *domain.Account) error {
	if !isEmailValid(acc.Email) {
		return fmt.Errorf("%w: invalid email format", domain.BadInput)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%w: encountered an error hashing password", err)
	}

	acc.Password = string(hash)
	err = a.store.Create(acc)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountService) Delete(id, accId int) error {
	if accId != id {
		return fmt.Errorf("%w: account with that id cannot be deleted", domain.Forbidden)
	}
	err := a.store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountService) FindByUsername(username string, accId int) (domain.Account, error) {
	acc, err := a.store.GetByUsername(username)
	if err != nil {
		return domain.Account{}, err
	}

	if acc.Id != accId {
		return domain.Account{}, fmt.Errorf("%w: account cannot be accessed with provided credentials", domain.Forbidden)
	}
	return acc, nil
}

var emailRegex = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
