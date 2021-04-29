package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoomService_FindRoom(t *testing.T) {
	store := new(mock.RoomStore)
	rs := NewRoomService(store)

	expectedRoom := domain.Room{
		RoomCode: "room1", Host: "Mathew",
	}
	store.On("GetByRoomCode", "room1").Return(expectedRoom, nil)
	room, err := rs.FindRoom("room1")

	assert.NoError(t, err)
	assert.EqualValues(t, expectedRoom, room)
	store.AssertExpectations(t)

	store.On("GetByRoomCode", "wrongCode").Return(domain.Room{}, errors.New("no room"))
	room, err = rs.FindRoom("wrongCode")
	assert.ErrorIs(t, err, errRoomNotFound)
	assert.EqualValues(t, domain.Room{}, room)
	store.AssertExpectations(t)
}

func TestRoomService_CreateRoom(t *testing.T) {
	store := new(mock.RoomStore)
	rs := NewRoomService(store)

	roomCreating := domain.Room{RoomCode: "room1", Host: "Mat"}
	store.On("Create", &roomCreating).Return(nil)
	err := rs.CreateRoom(&roomCreating)

	assert.NoError(t, err)
	store.AssertExpectations(t)

	wrongRoom := domain.Room{RoomCode: "duplicateCode", Host: "Ja"}
	store.On("Create", &wrongRoom).Return(errDuplicateRoom)
	err = rs.CreateRoom(&wrongRoom)

	assert.ErrorIs(t, err, errDuplicateRoom)
	store.AssertExpectations(t)

	err = rs.CreateRoom(&domain.Room{})
	assert.ErrorIs(t, err, errHostNotAssigned)
}

func TestRoomService_DeleteRoom(t *testing.T) {
	store := new(mock.RoomStore)
	rs := NewRoomService(store)

	validCode := "roomCode"
	store.On("Delete", validCode).Return(nil)
	err := rs.DeleteRoom(validCode)
	assert.NoError(t, err)

	store.On("Delete", "wrongCode").Return(errors.New("nothing deleted"))
	err = rs.DeleteRoom("wrongCode")
	assert.Error(t, err)
}
