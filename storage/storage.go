package storage

import (
	"github.com/DanielMolinaR/bank-project/model"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccounts() ([]*model.Account, error)
	GetAccountsFromCustomer(int) ([]*model.Account, error)
	GetAccountByID(int) (*model.Account, error)
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(*model.Account) error

	GetCustomers() ([]*model.Customer, error)
	GetCustomerByID(int) (*model.Customer, error)
	CreateCustomer(*model.Customer) error
	DeleteCustomer(int) error
	UpdateCustomer(*model.Customer) error
}
