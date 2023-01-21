package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	customer := NewCustomer("d", "mr", "example@gmail.com", "+31 123456789")
	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	createCustomerReq := &CreateCustomerRequest{}
	if err := json.NewDecoder(r.Body).Decode(createCustomerReq); err != nil {
		return err
	}

	customer := NewCustomer(
		createCustomerReq.FirstName,
		createCustomerReq.LastName,
		createCustomerReq.Email,
		createCustomerReq.PhoneNumber,
	)

	if err := s.store.CreateCustomer(customer); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
