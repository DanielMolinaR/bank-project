package model

import (
	"math/rand"
	"time"
)

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
