package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"golang.org/x/crypto/bcrypt"
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type accountService struct {
	store domain.AccountStore
}

func NewAccountService(store domain.AccountStore) domain.AccountService {
	return &accountService{store}
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

func (a *accountService) Delete(id int) error {
	err := a.store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

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
