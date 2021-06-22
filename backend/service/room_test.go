package service

import (
	"errors"
	"fmt"
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

	roomCreating := domain.Room{RoomCode: "room1", Host: "Mat", AccId: 1}
	store.On("Create", &roomCreating).Return(nil)
	err := rs.CreateRoom(&roomCreating)

	assert.NoError(t, err)
	store.AssertExpectations(t)

	wrongRoom := domain.Room{RoomCode: "duplicateCode", Host: "Ja", AccId: 1}
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

	room := domain.Room{
		RoomCode: "validCode",
		Host:     "Mat",
		AccId:    1,
	}

	store.On("Delete", "validCode").Return(nil)
	store.On("GetByRoomCode", "validCode").Return(room, nil)
	err := rs.DeleteRoom("validCode", 1)
	assert.NoError(t, err)

	store.On("GetByRoomCode", "wrongCode").Return(domain.Room{}, errors.New("no room"))
	err = rs.DeleteRoom("wrongCode", 1)
	assert.Error(t, err)

	err = rs.DeleteRoom("validCode", 3)
	assert.ErrorIs(t, err, domain.Forbidden)
}

func TestRoomService_AllRoomsWithId(t *testing.T) {
	store := new(mock.RoomStore)
	rs := NewRoomService(store)

	result := []domain.Room{
		{
			Host:     "Mat",
			RoomCode: "ACode",
			AccId:    1,
		},
		{
			Host:     "Mat",
			RoomCode: "AnotherCode",
			AccId:    1,
		},
	}
	store.On("FindAllRooms", 1).Return(result, nil)

	rooms, err := rs.AllRoomsWithId(1)
	assert.NoError(t, err)
	assert.EqualValues(t, result, rooms)

	store.On("FindAllRooms", 2).
		Return([]domain.Room{}, fmt.Errorf("%w: connection error", domain.Internal))
	_, err = rs.AllRoomsWithId(2)
	assert.Error(t, err)
}
