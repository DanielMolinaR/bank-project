package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielMolinaR/bank-project/model"
	"github.com/gorilla/mux"
)

type CreateCustomerRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (s *APIServer) handleCustomer(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetCustomerByID(w, r)
	case "POST":
		return s.handleCreateCustomer(w, r)
	case "DELETE":
		return s.handleDeleteCustomer(w, r)
	default:
		return fmt.Errorf("method not allowed %s ", r.Method)
	}
}

func (s *APIServer) handleGetCustomers(w http.ResponseWriter, r *http.Request) error {
	customers, err := s.store.GetCustomers()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, customers)
}

func (s *APIServer) handleGetCustomerByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	customer := model.NewCustomer("d", "mr", "example@gmail.com", "+31 123456789")
	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	createCustomerReq := &CreateCustomerRequest{}
	if err := json.NewDecoder(r.Body).Decode(createCustomerReq); err != nil {
		return err
	}

	customer := model.NewCustomer(
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
