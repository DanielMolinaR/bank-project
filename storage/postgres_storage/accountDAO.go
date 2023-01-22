package postgres_storage

import (
	"fmt"

	"github.com/DanielMolinaR/bank-project/model"
)

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT exists account (
		id 			serial primary key,
		customer_id int references customer,
		number 		int,
		balance		float,
		created_at 	timestamp
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
		account := new(model.Account)
		err := rows.Scan(
			&account.ID,
			&account.CustomerID,
			&account.Number,
			&account.Balance,
			&account.CreatedAt)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*model.Account, error) {
	return nil, nil
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
	return nil
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}
