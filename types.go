package main

import "math/rand"

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_nme"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Account struct {
	ID      int     `json:"id"`
	Client  *Person `json:"client_data"`
	Number  int64   `json:"number"`
	Balance int64   `json:"balance"`
}

func NewPerson(firstName, lastName, email, phoneNumber string) *Person {
	return &Person{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func NewAccount(firstName, lastName, email, phoneNumber string) *Account {
	return &Account{
		ID:     rand.Intn(10000),
		Client: NewPerson(firstName, lastName, email, phoneNumber),
		Number: int64(rand.Intn(100000)),
	}
}
