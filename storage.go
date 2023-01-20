package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccountByID(int) (*Account, error)
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	err := s.CreateCustomerTable()
	if err != nil {
		return err
	}
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateCustomerTable() error {
	query := `CREATE TABLE IF NOT exists customer (
		id 				integer primary key,
		first_name   	varchar(50),
		last_name    	varchar(50),
		email      		varchar(50)
		phone_number 	varchar(13)
	)`

	_, err := s.db.Exec(query)
	return err
}

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

func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
