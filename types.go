package main

import (
	"math/rand"
	"time"
)

type CreateCustomerRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func NewCustomer(firstName, lastName, email, phoneNumber string) *Customer {
	return &Customer{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

type CreateAccountRequest struct {
	CustomerID int `json:"customer_id"`
}

type Account struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Number     int64     `json:"number"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewAccount(customer_id int) *Account {
	return &Account{
		ID:         rand.Intn(10000),
		CustomerID: customer_id,
		Number:     int64(rand.Intn(100000)),
		CreatedAt:  time.Now().UTC(),
	}
}
