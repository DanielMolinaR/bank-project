package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type CreateAccountRequest struct {
	CustomerID string `json:"customer_id"`
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

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method not allowed %s ", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	account := NewAccount(1)
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}

	account := NewAccount(1)

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
