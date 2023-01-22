package postgres_storage

import (
	"database/sql"
	"fmt"

	"github.com/DanielMolinaR/bank-project/model"
)

func (s *PostgresStore) CreateCustomerTable() error {
	query := `CREATE TABLE IF NOT EXISTS customer (
		id 				SERIAL PRIMARY KEY,
		first_name   	VARCHAR(50),
		last_name    	VARCHAR(50),
		email      		VARCHAR(50),
		phone_number 	VARCHAR(13)
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetCustomers() ([]*model.Customer, error) {
	query := `SELECT * FROM customer`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	customers := []*model.Customer{}
	for rows.Next() {
		customer, err := ScanIntoCustomer(rows)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (s *PostgresStore) GetCustomerByID(id int) (*model.Customer, error) {
	query := "SELECT * FROM customer WHERE ID = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanIntoCustomer(rows)
	}

	return nil, fmt.Errorf("customer %d not found", id)
}

func (s *PostgresStore) CreateCustomer(customer *model.Customer) error {
	query := `
	insert into customer 
	(first_name, last_name, email, phone_number)
	values 
	($1, $2, $3, $4)`

	resp, err := s.db.Query(
		query,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.PhoneNumber)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) DeleteCustomer(id int) error {
	query := "DELETE FROM customer WHERE id = $1"

	_, err := s.db.Query(query, id)

	return err
}

func (s *PostgresStore) UpdateCustomer(*model.Customer) error {
	return nil
}

func ScanIntoCustomer(rows *sql.Rows) (*model.Customer, error) {
	customer := &model.Customer{}
	err := rows.Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.PhoneNumber)

	return customer, err
}
