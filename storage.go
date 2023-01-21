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
	GetCustomerByID(int) (*Customer, error)
	CreateCustomer(*Customer) error
	DeleteCustomer(int) error
	UpdateCustomer(*Customer) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
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
