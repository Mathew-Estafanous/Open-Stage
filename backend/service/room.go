package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"math/rand"
	"time"
)

type roomService struct {
	rStore domain.RoomStore
}

func NewRoomService(rs domain.RoomStore) domain.RoomService {
	return &roomService{
		rStore: rs,
	}
}

func (r *roomService) CreateRoom(room *domain.Room) error {
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

func (r *roomService) FindRoom(roomCode string) (domain.Room, error) {
	room, err := r.rStore.GetByRoomCode(roomCode)
	if err != nil {
		return domain.Room{}, errRoomNotFound
	}
	return room, nil
}

func (r *roomService) DeleteRoom(code string) error {
	err := r.rStore.Delete(code)
	if err != nil {
		return errRoomNotDeleted
	}
	return nil
}

func (r *roomService) generateValidCode() string {
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
	errHostNotAssigned = domain.ApiError{Msg: "A host has not be assigned to a room", Typ: domain.BadInput}
	errDuplicateRoom   = domain.ApiError{Msg: "Room could not be created since the room code is taken.", Typ: domain.Conflict}
	errRoomNotFound    = domain.ApiError{Msg: "Room was not found with given code", Typ: domain.NotFound}
	errRoomNotDeleted  = domain.ApiError{Msg: "There was an error while trying to delete the room.", Typ: domain.Internal}
)
