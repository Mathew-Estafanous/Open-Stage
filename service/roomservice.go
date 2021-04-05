package service

import (
	"errors"
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
		return ErrHostNotAssigned
	}

	if room.RoomCode == "" {
		room.RoomCode = r.generateValidCode()
	}

	err := r.rStore.Create(room)
	if err != nil {
		return ErrDuplicateRoom
	}
	return nil
}

func (r *roomService) FindRoom(roomCode string) (domain.Room, error) {
	room, err := r.rStore.GetByRoomCode(roomCode)
	if err != nil {
		return domain.Room{}, ErrRoomNotFound
	}
	return room, nil
}

func (r *roomService) DeleteRoom(code string) error {
	err := r.rStore.Delete(code)
	if err != nil {
		return ErrRoomNotDeleted
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
	ErrHostNotAssigned = errors.New("a host has not be assigned to a room")
	ErrDuplicateRoom   = errors.New("room could not be created with duplicate room code")
	ErrRoomNotFound    = errors.New("room was not found with given code")
	ErrRoomNotDeleted  = errors.New("a room with that code was not found")
)
