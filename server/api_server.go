package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/DanielMolinaR/bank-project/model"
	"github.com/DanielMolinaR/bank-project/storage"
	"github.com/golang-jwt/jwt/v4"
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

	router.HandleFunc("/customers", withJWTAuth(makeHTTPHandleFunc(s.handleGetCustomers), s.store))
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

func createJWT(customer *model.Customer) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":  15000,
		"customerID": customer.ID,
		"firstName":  customer.FirstName,
		"lastName":   customer.LastName,
		"email":      customer.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, apiError{Error: "Permission denied"})
}

func withJWTAuth(handleFunc http.HandlerFunc, s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling JWT auth")

		reqToken := r.Header.Get("Authorization")

		tokenStr := strings.Split(reqToken, "Bearer ")

		token, err := validateJWT(strings.Join(tokenStr, ""))

		if err != nil {
			fmt.Println(err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)

		_, err = s.GetCustomerByID(int(claims["customerID"].(float64)))

		if err != nil {
			fmt.Println(err)
			permissionDenied(w)
			return
		}

		handleFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

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
