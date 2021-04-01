package database

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetByRoomCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mockRoom := domain.Room{
		RoomId: 1, RoomCode: "wantedCode", Host: "Mathew",
	}

	row := sqlmock.NewRows([]string{"room_id", "host", "room_code"}).
		AddRow(mockRoom.RoomId, mockRoom.Host, mockRoom.RoomCode)

	query := "SELECT room_id, host, room_code FROM rooms WHERE room_code = ?"
	mock.ExpectQuery(query).WithArgs("wantedCode").WillReturnRows(row)

	m := NewMySQLRoomStore(db)
	room, err := m.GetByRoomCode("wantedCode")

	assert.NoError(t, err)
	assert.EqualValues(t, mockRoom, room)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	room := &domain.Room{
		RoomCode: "jrhigh",
		Host:     "Mathew",
	}

	insertQuery := "INSERT INTO rooms"
	mock.ExpectExec(insertQuery).WithArgs(room.Host, room.RoomCode).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := NewMySQLRoomStore(db)
	err = m.Create(room)
	assert.NoError(t, err)
}
