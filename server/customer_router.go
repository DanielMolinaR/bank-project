package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielMolinaR/bank-project/model"
)

type CreateCustomerRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (s *APIServer) handleCustomerById(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetCustomerByID(w, r)
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
	id, err := getID(r)

	if err != nil {
		return err
	}

	customer, err := s.store.GetCustomerByID(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleCreateCustomer(w http.ResponseWriter, r *http.Request) error {
	createCustomerReq := &CreateCustomerRequest{}
	if err := json.NewDecoder(r.Body).Decode(createCustomerReq); err != nil {
		return err
	}
	defer r.Body.Close()

	customer := model.NewCustomer(
		createCustomerReq.FirstName,
		createCustomerReq.LastName,
		createCustomerReq.Email,
		createCustomerReq.PhoneNumber,
	)

	tokenStr, err := createJWT(customer)

	if err != nil {
		return err
	}

	fmt.Println(tokenStr)

	if err := s.store.CreateCustomer(customer); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, customer)
}

func (s *APIServer) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)

	if err != nil {
		return err
	}

	if err = s.store.DeleteCustomer(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}
