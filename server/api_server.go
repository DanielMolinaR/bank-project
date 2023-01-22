package server

import (
	"encoding/json"
	"log"
	"net/http"

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

	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleGetAccounts))

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/customers", makeHTTPHandleFunc(s.handleGetCustomers))

	router.HandleFunc("/customer", makeHTTPHandleFunc(s.handleCustomer))

	router.HandleFunc("/customer/{id}", makeHTTPHandleFunc(s.handleCustomer))

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
	Error string
}

// Decorates apiFunc so it looks like a HandleFunc function from gorilla mux
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}
