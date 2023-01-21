package main

import "fmt"

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT exists account (
		id 			serial primary key,
		number 		serial,
		customer_id int references customer,
		balance		float,
		created_at 	timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `
	insert into account 
	(number, customer_id, balance, created_at)
	values 
	($1, $2, $3, $4)`

	resp, err := s.db.Query(
		query,
		acc.Number,
		acc.CustomerID,
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

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
