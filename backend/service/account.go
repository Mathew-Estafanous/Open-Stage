package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type accountService struct {
	key   string
	store domain.AccountStore
}

type AccountClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAccountService(aStore domain.AccountStore) domain.AccountService {
	return &accountService{
		store: aStore,
		key:   os.Getenv("SECRET_KEY"),
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

func (a *accountService) Authenticate(acc domain.Account) (domain.AuthToken, error) {
	found, err := a.store.GetByUsername(acc.Username)
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: could not find with that username", domain.Unauthorized)
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(acc.Password))
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: the password did not match", domain.Unauthorized)
	}

	tk, err := createToken(found.Username, strconv.Itoa(found.Id), a.key)
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: we were unable to issue a token", domain.Internal)
	}
	return tk, nil
}

func createToken(username, id , key string) (domain.AuthToken, error) {
	exp := time.Now().Add(time.Minute * 15).Unix()
	accessClaim := AccountClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Audience: "access",
			Subject: id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	access, err := token.SignedString([]byte(key))
	if err != nil {
		return domain.AuthToken{}, err
	}

	exp = time.Now().Add(time.Hour * 168).Unix()
	refreshClaim := AccountClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Audience: "refresh",
			Subject: id,
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refresh, err := token.SignedString([]byte(key))
	if err != nil {
		return domain.AuthToken{}, err
	}

	tk := domain.AuthToken{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	return tk, nil
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
