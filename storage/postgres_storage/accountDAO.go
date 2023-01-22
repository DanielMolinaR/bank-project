package postgres_storage

import (
	"database/sql"
	"fmt"

	"github.com/DanielMolinaR/bank-project/model"
)

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id 			SERIAL PRIMARY KEY,
		customer_id INT references customer(id) ON DELETE CASCADE,
		number 		INT,
		balance		FLOAT,
		created_at 	TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetAccounts() ([]*model.Account, error) {
	query := `SELECT * FROM account`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}
	for rows.Next() {
		account, err := ScanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountsFromCustomer(id int) ([]*model.Account, error) {
	query := `SELECT * FROM account WHERE customer_id = $1`

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}
	for rows.Next() {
		account, err := ScanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*model.Account, error) {
	query := "SELECT * FROM account WHERE ID = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) CreateAccount(acc *model.Account) error {
	query := `
	insert into account 
	(customer_id, number, balance, created_at)
	values 
	($1, $2, $3, $4)`

	resp, err := s.db.Query(
		query,
		acc.CustomerID,
		acc.Number,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := "DELETE FROM account WHERE id = $1"

	_, err := s.db.Query(query, id)

	return err
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}

func ScanIntoAccount(rows *sql.Rows) (*model.Account, error) {
	account := &model.Account{}
	err := rows.Scan(
		&account.ID,
		&account.CustomerID,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
