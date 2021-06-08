package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type authService struct {
	store domain.AccountStore
	key string
}

func NewAuthService(acc domain.AccountStore) domain.AuthService {
	return authService{
		store: acc,
		key:   os.Getenv("SECRET_KEY"),
	}
}

func (a authService) Authenticate(username, password string) (domain.AuthToken, error) {
	found, err := a.store.GetByUsername(username)
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: could not find with that username", domain.Unauthorized)
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(password))
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: the password did not match", domain.Unauthorized)
	}

	tk, err := createToken(found.Username, strconv.Itoa(found.Id), a.key)
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: we were unable to issue a token", domain.Internal)
	}
	return tk, nil
}

type AccountClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createToken(username, id, key string) (domain.AuthToken, error) {
	exp := time.Now().Add(time.Minute * 15).Unix()
	accessClaim := AccountClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Audience:  "access",
			Subject:   id,
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
			Audience:  "refresh",
			Subject:   id,
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
