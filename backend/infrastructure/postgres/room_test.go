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
		RoomCode: "wantedCode", Host: "Mathew", AccId: 1,
	}

	row := sqlmock.NewRows([]string{"host", "room_code", "fk_account_id"}).
		AddRow(mockRoom.Host, mockRoom.RoomCode, mockRoom.AccId)

	query := "SELECT host, room_code, fk_account_id FROM rooms WHERE room_code = ?"
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
		AccId:    1,
	}

	insertQuery := "INSERT INTO rooms"
	mock.ExpectExec(insertQuery).WithArgs(room.Host, room.RoomCode, room.AccId).
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

func TestPostgresRoomStore_FindAllRooms(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	result := []domain.Room{
		{
			Host:     "Mat",
			RoomCode: "ARoomCode",
			AccId:    1,
		},
	}

	mock.ExpectQuery("SELECT host, room_code, fk_account_id FROM rooms").WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"host", "room_code", "fk_account_id"}).
			AddRow("Mat", "ARoomCode", 1))

	m := NewRoomStore(db)
	rooms, err := m.FindAllRooms(1)
	assert.NoError(t, err)
	assert.EqualValues(t, result, rooms)
}
