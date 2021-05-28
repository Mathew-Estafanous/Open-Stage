package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net"
	"regexp"
	"strings"
	"time"
)

type accountService struct {
	sign jwt.SigningMethod
	store domain.AccountStore
}

type AccountClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAccountService(aStore domain.AccountStore) domain.AccountService {
	return &accountService{
		store: aStore,
		sign: jwt.SigningMethodHS256,
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

func (a *accountService) Delete(id int) error {
	err := a.store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountService) Authenticate(acc domain.Account) (string, error) {
	found, err := a.store.GetByUsername(acc.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(acc.Password))
	if err != nil {
		return "", fmt.Errorf("%w: the password did not match", domain.BadInput)
	}

	exp := time.Now().Add(time.Minute * 15).Unix()
	claim := AccountClaims{
		found.Username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer: "server",
		},
	}
	token := jwt.NewWithClaims(a.sign, claim)
	signedToken, err := token.SignedString([]byte("SECRETKEY"))
	if err != nil {
		return "", fmt.Errorf("%w: encountered an eror while signing the token", err)
	}
	return signedToken, nil
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
