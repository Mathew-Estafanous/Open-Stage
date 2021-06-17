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
	row := p.db.QueryRow("SELECT host, room_code, fk_account_id FROM rooms WHERE room_code = $1", code)

	var room domain.Room
	err := row.Scan(&room.Host, &room.RoomCode, &room.AccId)
	if err != nil {
		return domain.Room{}, err
	}
	return room, err
}

func (p *postgresRoomStore) Create(room *domain.Room) error {
	_, err := p.db.Exec("INSERT INTO rooms (host, room_code, fk_account_id) VALUES ($1, $2, $3)",
		room.Host, room.RoomCode, room.AccId)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresRoomStore) Delete(roomCode string) error {
	r, err := p.db.Exec("DELETE FROM rooms WHERE room_code = $1", roomCode)
	if err != nil {
		return err
	}

	if a, _ := r.RowsAffected(); a == 0 {
		return errors.New("no rows were deleted")
	}
	return nil
}

func (p *postgresRoomStore) FindAllRooms(accId int) ([]domain.Room, error) {
	rows, err := p.db.Query("SELECT host, room_code, fk_account_id FROM rooms WHERE fk_account_id = $1", accId)
	if err != nil {
		return nil, err
	}

	rooms := make([]domain.Room, 0)
	for rows.Next() {
		var r domain.Room
		err = rows.Scan(&r.Host, &r.RoomCode, &r.AccId)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, nil
}

