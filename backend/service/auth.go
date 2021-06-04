package service

import "github.com/Mathew-Estafanous/Open-Stage/domain"

type authService struct {
	rStore domain.RoomStore
}

func NewAuthService(r domain.RoomStore) domain.AuthService {
	return authService{
		rStore: r,
	}
}

func (a authService) OwnsRoom(code string, accId int) (bool, error) {
	room, err := a.rStore.GetByRoomCode(code)
	if err != nil {
		return false, err
	}

	return room.AccId == accId, nil
}

