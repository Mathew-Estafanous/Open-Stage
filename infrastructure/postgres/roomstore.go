package postgres

import (
	"database/sql"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

// mySQLRoomStore is a MySQL implementation of the RoomStore
// interface and the associated contracts that provided.
type mySQLRoomStore struct {
	db *sql.DB
}

func NewRoomStore(db *sql.DB) domain.RoomStore {
	return &mySQLRoomStore{db}
}

func (m *mySQLRoomStore) GetByRoomCode(code string) (domain.Room, error) {
	row := m.db.QueryRow("SELECT room_id, host, room_code FROM rooms WHERE room_code = $1", code)

	var room domain.Room
	err := row.Scan(&room.RoomId, &room.Host, &room.RoomCode)
	if err != nil {
		return domain.Room{}, err
	}
	return room, err
}

func (m *mySQLRoomStore) Create(room *domain.Room) error {
	r, err := m.db.Exec("INSERT INTO rooms (host, room_code) VALUES ($1, $2)",
		room.Host, room.RoomCode)
	if err != nil {
		return err
	}
	id, _ := r.LastInsertId()
	room.RoomId = int(id)
	return nil
}

func (m *mySQLRoomStore) Delete(roomCode string) error {
	r, err := m.db.Exec("DELETE FROM rooms WHERE room_code = $3", roomCode)
	if err != nil {
		return err
	}

	if a, _ := r.RowsAffected(); a == 0 {
		return errors.New("no rows were deleted")
	}
	return nil
}
