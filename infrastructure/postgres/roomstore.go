package postgres

import (
	"database/sql"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

// postgresRoomStore is a database implementation of the RoomStore
// interface and provides the functionality for all the required operations.
type postgresRoomStore struct {
	db *sql.DB
}

func NewRoomStore(db *sql.DB) domain.RoomStore {
	return &postgresRoomStore{db}
}

func (p *postgresRoomStore) GetByRoomCode(code string) (domain.Room, error) {
	row := p.db.QueryRow("SELECT room_id, host, room_code FROM rooms WHERE room_code = $1", code)

	var room domain.Room
	err := row.Scan(&room.RoomId, &room.Host, &room.RoomCode)
	if err != nil {
		return domain.Room{}, err
	}
	return room, err
}

func (p *postgresRoomStore) Create(room *domain.Room) error {
	r, err := p.db.Exec("INSERT INTO rooms (host, room_code) VALUES ($1, $2)",
		room.Host, room.RoomCode)
	if err != nil {
		return err
	}
	id, _ := r.LastInsertId()
	room.RoomId = int(id)
	return nil
}

func (p *postgresRoomStore) Delete(roomCode string) error {
	r, err := p.db.Exec("DELETE FROM rooms WHERE room_code = $3", roomCode)
	if err != nil {
		return err
	}

	if a, _ := r.RowsAffected(); a == 0 {
		return errors.New("no rows were deleted")
	}
	return nil
}
