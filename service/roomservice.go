package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"math/rand"
	"time"
)

type roomService struct {
	roomStore domain.RoomStore
}

func NewRoomService(rs domain.RoomStore) domain.RoomService {
	return &roomService{
		roomStore: rs,
	}
}

func (r *roomService) CreateRoom(room *domain.Room) error {
	if room.Host == "" {
		return ErrHostNotAssigned
	}

	if room.RoomCode == "" {
		room.RoomCode = r.generateValidCode()
	}

	err := r.roomStore.Create(room)
	if err != nil {
		return ErrRoomNotCreated
	}
	return nil
}

func (r *roomService) FindRoom(roomCode string) (domain.Room, error) {
	room, err := r.roomStore.GetByRoomCode(roomCode)
	if err != nil {
		return domain.Room{}, ErrRoomNotFound
	}
	return room, nil
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

	_, err := r.roomStore.GetByRoomCode(sequence)
	if err == nil {
		return r.generateValidCode()
	}
	return sequence
}

var (
	ErrHostNotAssigned = errors.New("a host has not be assigned to a room")
	ErrRoomNotCreated  = errors.New("room could not be created with duplicate room code")
	ErrRoomNotFound    = errors.New("room was not found with given code")
)
