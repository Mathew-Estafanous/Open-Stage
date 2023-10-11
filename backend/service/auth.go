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
	cache domain.AuthCache
	key   string
}

func NewAuthService(acc domain.AccountStore, cache domain.AuthCache) domain.AuthService {
	return authService{
		store: acc,
		cache: cache,
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

func (a authService) Refresh(refreshTkn string) (domain.AuthToken, error) {
	blacklisted, err := a.cache.Contains(refreshTkn)
	if err != nil {
		return domain.AuthToken{}, err
	}
	if blacklisted {
		return domain.AuthToken{}, fmt.Errorf("%w: Provided token has been blacklisted", domain.Unauthorized)
	}

	tkn, err := jwt.Parse(refreshTkn, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.key), nil
	})
	if err != nil {
		return domain.AuthToken{}, fmt.Errorf("%w: invalid refresh token", domain.Unauthorized)
	}

	c := tkn.Claims.(jwt.MapClaims)
	if c["aud"] != "refresh" {
		return domain.AuthToken{}, fmt.Errorf("%w: invalid refresh token", domain.Unauthorized)
	}
	id := c["sub"].(string)
	username := c["username"].(string)
	return createToken(username, id, a.key)
}

func (a authService) Invalidate(token domain.AuthToken) error {
	err := a.cache.Store(token.AccessToken)
	if err != nil {
		return err
	}
	err = a.cache.Store(token.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}

type AccountClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createToken(username, id, key string) (domain.AuthToken, error) {
	exp := time.Now().Add(domain.AccessTokenTimeout).Unix()
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

	exp = time.Now().Add(domain.RefreshTokenTimeout).Unix()
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
