package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresRoomStore_GetByRoomCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mockRoom := domain.Room{
		RoomCode: "wantedCode", Host: "Mathew",
	}

	row := sqlmock.NewRows([]string{"host", "room_code"}).
		AddRow(mockRoom.Host, mockRoom.RoomCode)

	query := "SELECT host, room_code FROM rooms WHERE room_code = ?"
	mock.ExpectQuery(query).WithArgs("wantedCode").WillReturnRows(row)

	m := NewRoomStore(db)
	room, err := m.GetByRoomCode("wantedCode")

	assert.NoError(t, err)
	assert.EqualValues(t, mockRoom, room)
}

func TestPostgresRoomStore_Create(t *testing.T) {
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

	m := NewRoomStore(db)
	err = m.Create(room)
	assert.NoError(t, err)
}

func TestPostgresRoomStore_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	code := "roomCode"
	deleteQuery := "DELETE FROM rooms"
	mock.ExpectExec(deleteQuery).WithArgs(code).
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := NewRoomStore(db)
	err = m.Delete(code)
	assert.NoError(t, err)

	mock.ExpectExec(deleteQuery).WithArgs("wrongCode").
		WillReturnResult(sqlmock.NewResult(0, 0))
	err = m.Delete("wrongCode")
	assert.Error(t, err)
}
