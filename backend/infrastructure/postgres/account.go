package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/lib/pq"
)

type postgresAccountStore struct {
	db *sql.DB
}

func NewAccountStore(db *sql.DB) domain.AccountStore {
	return &postgresAccountStore{db}
}

func (p *postgresAccountStore) Create(acc *domain.Account) error {
	r, err := p.db.Query("INSERT INTO accounts (name, username, password, email) VALUES ($1, $2, $3, $4) RETURNING id",
		acc.Name, acc.Username, acc.Password, acc.Email)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return fmt.Errorf("%w: duplicate username key value", domain.Conflict)
			}
		}
		return err
	}
	r.Next()
	err = r.Scan(&acc.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresAccountStore) GetByUsername(username string) (domain.Account, error) {
	r, err := p.db.Query("SELECT id, name, username, password, email FROM accounts WHERE username = $1", username)
	if err != nil {
		return domain.Account{}, err
	}

	if !r.Next() {
		return domain.Account{}, fmt.Errorf("%w: could not find an account", domain.NotFound)
	}

	var account domain.Account
	err = r.Scan(&account.Id, &account.Name, &account.Username, &account.Password, &account.Email)
	if err != nil {
		return domain.Account{}, err
	}
	return account, err
}

func (p *postgresAccountStore) Delete(id int) error {
	_, err := p.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
