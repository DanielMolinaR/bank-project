package main

import "fmt"

func (s *PostgresStore) CreateCustomerTable() error {
	query := `CREATE TABLE IF NOT exists customer (
		id 				serial primary key,
		first_name   	varchar(50),
		last_name    	varchar(50),
		email      		varchar(50),
		phone_number 	varchar(13)
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetCustomerByID(id int) (*Customer, error) {
	return nil, nil
}

func (s *PostgresStore) CreateCustomer(customer *Customer) error {
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
	return nil
}

func (s *PostgresStore) UpdateCustomer(*Customer) error {
	return nil
}
