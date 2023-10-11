package domain

import "time"

const (
	AccessTokenTimeout  = time.Minute * 15
	RefreshTokenTimeout = time.Hour * 168
)

// AuthToken contains both the access and refresh tokens after
// a user has successfully authenticated.
//
// swagger:model authToken
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}

type AuthService interface {
	Authenticate(username, password string) (AuthToken, error)
	Refresh(refreshTkn string) (AuthToken, error)
	Invalidate(token AuthToken) error
}

type AuthCache interface {
	Contains(tkn string) (bool, error)
	Store(tkn string) error
}
