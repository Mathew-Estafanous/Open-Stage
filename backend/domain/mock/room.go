package mock

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/mock"
)

// RoomService is a mock struct that is used in unit tests
type RoomService struct {
	mock.Mock
}

func (m *RoomService) CreateRoom(room *domain.Room) error {
	ret := m.Called(room)
	return ret.Error(0)
}

func (m *RoomService) FindRoom(roomCode string) (domain.Room, error) {
	ret := m.Called(roomCode)
	return ret.Get(0).(domain.Room), ret.Error(1)
}

func (m *RoomService) DeleteRoom(code string, accId int) error {
	ret := m.Called(code, accId)
	return ret.Error(0)
}

func (m *RoomService) AllRoomsWithId(accId int) ([]domain.Room, error) {
	ret := m.Called(accId)
	return ret.Get(0).([]domain.Room), ret.Error(1)
}

// RoomStore is a mock struct that is used in unit tests.
type RoomStore struct {
	mock.Mock
}

func (m *RoomStore) GetByRoomCode(code string) (domain.Room, error) {
	ret := m.Called(code)
	return ret.Get(0).(domain.Room), ret.Error(1)
}

func (m *RoomStore) Create(room *domain.Room) error {
	ret := m.Called(room)
	return ret.Error(0)
}

func (m *RoomStore) Delete(code string) error {
	ret := m.Called(code)
	return ret.Error(0)
}

func (m *RoomStore) FindAllRooms(accId int) ([]domain.Room, error) {
	ret := m.Called(accId)
	return ret.Get(0).([]domain.Room), ret.Error(1)
}
