package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID                int    `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	EncryptedPassword string `json:"encrypted_password"`
}

func NewCustomer(firstName, lastName, email, phoneNumber, password string) (*Customer, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Customer{
		FirstName:         firstName,
		LastName:          lastName,
		Email:             email,
		PhoneNumber:       phoneNumber,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}

func (c *Customer) ValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(c.EncryptedPassword), []byte(password)) == nil
}
