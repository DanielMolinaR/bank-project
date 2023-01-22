package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DanielMolinaR/bank-project/storage"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewApiServer(listenAdrr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAdrr,
		store:      store,
	}
}

func (s APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", makeHTTPHandleFunc(s.handleGetCustomers))
	router.HandleFunc("/customer", makeHTTPHandleFunc(s.handleCreateCustomer))
	router.HandleFunc("/customer/{id}", makeHTTPHandleFunc(s.handleCustomerById))

	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleGetAccounts))
	router.HandleFunc("/customer-accounts/{id}", makeHTTPHandleFunc(s.handleGetAccountsFromCustomer))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleCreateAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccountById))

	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.HandleTransfer))

	log.Println("JSON Api Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

// Decorates apiFunc so it looks like a HandleFunc function from gorilla mux
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
