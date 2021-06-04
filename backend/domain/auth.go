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
	OwnsRoom(code string, accId int) (bool, error)
}
