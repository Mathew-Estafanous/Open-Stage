package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRoomStore struct {
	mock.Mock
}

func (m *mockRoomStore) GetByRoomCode(code string) (domain.Room, error) {
	ret := m.Called(code)
	return ret.Get(0).(domain.Room), ret.Error(1)
}

func (m *mockRoomStore) Create(room *domain.Room) error {
	ret := m.Called(room)
	return ret.Error(0)
}

func (m *mockRoomStore) Delete(code string) error {
	ret := m.Called(code)
	return ret.Error(0)
}

func TestRoomService_FindRoom(t *testing.T) {
	store := new(mockRoomStore)
	rs := NewRoomService(store)

	expectedRoom := domain.Room{
		RoomId: 1, RoomCode: "room1", Host: "Mathew",
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
	store := new(mockRoomStore)
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
	store := new(mockRoomStore)
	rs := NewRoomService(store)

	validCode := "roomCode"
	store.On("Delete", validCode).Return(nil)
	err := rs.DeleteRoom(validCode)
	assert.NoError(t, err)

	store.On("Delete", "wrongCode").Return(errors.New("nothing deleted"))
	err = rs.DeleteRoom("wrongCode")
	assert.Error(t, err)
}
