package database

import (
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

// mySQLRoomRepository is a MySQL implementation of the RoomRepository
// interface and the associated contracts that provided.
type mySQLRoomRepository struct {
	db *sql.DB
}

func NewMySQLRoomRepository(db *sql.DB) domain.RoomRepository {
	return &mySQLRoomRepository{db}
}

func (m *mySQLRoomRepository) GetByCode(code string) (domain.Room, error) {
	row := m.db.QueryRow("SELECT room_id, host, room_code FROM rooms WHERE room_code = ?", code)

	var room domain.Room
	err := row.Scan(&room.RoomId, &room.Host)
	if err != nil {
		return domain.Room{}, err
	}
	return room, err
}

func (m *mySQLRoomRepository) Create(room *domain.Room) error {
	_, err := m.db.Exec("INSERT INTO rooms (host, room_code) VALUES (?, ?);",
		room.Host, room.RoomCode)
	if err != nil {
		return err
	}
	return nil
}
