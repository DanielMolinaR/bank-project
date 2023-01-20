package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_nme"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func NewCustomer(firstName, lastName, email, phoneNumber string) *Customer {
	return &Customer{
		ID:          rand.Intn(100000),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

func (s *APIServer) handleCustomer(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetCustomer(w, r)
	case "POST":
		return s.handleCreateCustomer(w, r)
	case "DELETE":
		return s.handleDeleteCustomer(w, r)
	default:
		return fmt.Errorf("method not allowed %s ", r.Method)
	}
}

func (s *APIServer) handleGetCustomer(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	account := NewAccount(1)
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	customer := NewCustomer("d", "mr", "example@gmail.com", "+31 123456789")

	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
