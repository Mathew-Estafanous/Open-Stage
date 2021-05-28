package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresAccountStore_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	acc := &domain.Account{
		Name: "Mathew",
		Username: "MatMat277",
		Password: "ThisIsAHashedPassword",
		Email: "mathew@gmail.com",
	}

	insertQuery := `INSERT INTO accounts (name, username, password, email)
						VALUES ($1, $2, $3, $4) RETURNING id`
	mock.ExpectQuery(insertQuery).
		WithArgs(acc.Name, acc.Username, acc.Password, acc.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	aStore := NewAccountStore(db)
	err = aStore.Create(acc)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, acc.Id)
}

func TestPostgresAccountStore_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	id := 1
	mock.ExpectExec("DELETE FROM accounts").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	aStore := NewAccountStore(db)
	err = aStore.Delete(id)
	assert.NoError(t, err)
}

func TestPostgresAccountStore_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	account := domain.Account{
		Id: 1,
		Name: "Mathew",
		Username: "MatMat277",
		Password: "ThisIsAHashedPassword",
		Email: "mathew@gmail.com",
	}

	row := sqlmock.NewRows([]string{"id", "name", "username", "password", "email"}).
		AddRow(account.Id, account.Name, account.Username, account.Password, account.Email)

	selectQuery := `SELECT id, name, username, password, email 
						FROM accounts WHERE username = $1`
	mock.ExpectQuery(selectQuery).WithArgs(account.Username).WillReturnRows(row)

	aStore := NewAccountStore(db)
	acc, err := aStore.GetByUsername(account.Username)
	assert.NoError(t, err)
	assert.EqualValues(t, account, acc)
}