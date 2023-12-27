package service

import (
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"math/rand"
	"time"
)

type RoomService struct {
	rStore domain.RoomStore
}

func NewRoomService(rs domain.RoomStore) *RoomService {
	return &RoomService{
		rStore: rs,
	}
}

func (r *RoomService) CreateRoom(room *domain.Room) error {
	if room.Host == "" {
		return errHostNotAssigned
	}

	if room.RoomCode == "" {
		room.RoomCode = r.generateValidCode()
	}

	err := r.rStore.Create(room)
	if err != nil {
		return errDuplicateRoom
	}
	return nil
}

func (r *RoomService) FindRoom(roomCode string) (domain.Room, error) {
	room, err := r.rStore.GetByRoomCode(roomCode)
	if err != nil {
		return domain.Room{}, errRoomNotFound
	}
	return room, nil
}

func (r *RoomService) DeleteRoom(code string, accId int) error {
	room, err := r.FindRoom(code)
	if err != nil {
		return err
	}

	if room.AccId != accId {
		return errDoesNotOwn
	}

	err = r.rStore.Delete(code)
	if err != nil {
		return errRoomNotDeleted
	}
	return nil
}

func (r *RoomService) AllRoomsWithId(accId int) ([]domain.Room, error) {
	rooms, err := r.rStore.FindAllRooms(accId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomService) generateValidCode() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	sequence := string(b)

	_, err := r.rStore.GetByRoomCode(sequence)
	if err == nil {
		return r.generateValidCode()
	}
	return sequence
}

var (
	errHostNotAssigned = fmt.Errorf("%w: a host has not be assigned to a room", domain.BadInput)
	errDuplicateRoom   = fmt.Errorf("%w: room could not be created since the room code is taken", domain.Conflict)
	errRoomNotFound    = fmt.Errorf("%w: room was not found with given code", domain.NotFound)
	errRoomNotDeleted  = fmt.Errorf("%w: we encountered an issue when trying to delete your room", domain.Internal)
	errDoesNotOwn      = fmt.Errorf("%w: account does not own room", domain.Forbidden)
)
