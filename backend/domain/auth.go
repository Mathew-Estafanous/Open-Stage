package domain

// AuthToken contains both the access and refresh tokens after
// a user has successfully authenticated.
//
// swagger:model authToken
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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
