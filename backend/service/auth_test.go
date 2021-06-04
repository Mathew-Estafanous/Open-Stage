package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_OwnsRoom(t *testing.T) {
	rStore := new(mock.RoomStore)
	auth := NewAuthService(rStore)
	rStore.On("GetByRoomCode", "validRoom").Return(domain.Room{
		RoomCode: "validRoom",
		AccId: 1,
	}, nil)

	doesOwn, err := auth.OwnsRoom("validRoom", 1)
	assert.NoError(t, err)
	assert.EqualValues(t, true, doesOwn)

	rStore.On("GetByRoomCode", "notOwnedRoom").Return(domain.Room{
		RoomCode: "notOwnedRoom",
		AccId: 2,
	}, nil)

	doesOwn, err = auth.OwnsRoom("notOwnedRoom", 1)
	assert.NoError(t, err)
	assert.EqualValues(t, false, doesOwn)
}
