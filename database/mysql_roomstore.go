package database

import (
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

// mySQLRoomStore is a MySQL implementation of the RoomStore
// interface and the associated contracts that provided.
type mySQLRoomStore struct {
	db *sql.DB
}

func NewMySQLRoomRepository(db *sql.DB) domain.RoomStore {
	return &mySQLRoomStore{db}
}

func (m *mySQLRoomStore) GetByRoomCode(code string) (domain.Room, error) {
	row := m.db.QueryRow("SELECT room_id, host, room_code FROM rooms WHERE room_code = ?", code)

	var room domain.Room
	err := row.Scan(&room.RoomId, &room.Host, &room.RoomCode)
	if err != nil {
		return domain.Room{}, err
	}
	return room, err
}

func (m *mySQLRoomStore) Create(room *domain.Room) error {
	_, err := m.db.Exec("INSERT INTO rooms (host, room_code) VALUES (?, ?);",
		room.Host, room.RoomCode)
	if err != nil {
		return err
	}
	return nil
}