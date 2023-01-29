package seed

import (
	"log"

	"github.com/DanielMolinaR/bank-project/model"
	"github.com/DanielMolinaR/bank-project/storage"
)

func SeedDB(s storage.Storage) {
	seedCustomers(s)
	seedAccounts(s)
}

func seedCustomer(fname, lname, email, phoneNumber, password string, store storage.Storage) *model.Customer {
	customer, err := model.NewCustomer(fname, lname, email, phoneNumber, password)
	if err != nil {
		log.Fatal(err)
	}

	if err = store.CreateCustomer(customer); err != nil {
		log.Fatal(err)
	}

	return customer
}

func seedCustomers(s storage.Storage) {
	var customers [3]map[string]string
	customers[0] = map[string]string{
		"fname":       "Bob",
		"lname":       "Ross",
		"email":       "bob.ross@gmail.com",
		"phoneNumber": "+31 123456789",
		"password":    "BobRoss9999",
	}
	customers[1] = map[string]string{
		"fname":       "Jeff",
		"lname":       "Mills",
		"email":       "jeff.mills@gmail.com",
		"phoneNumber": "+31 123456789",
		"password":    "JeffMills9999",
	}
	customers[2] = map[string]string{
		"fname":       "Seb",
		"lname":       "Wilder",
		"email":       "Seb.Wilder@gmail.com",
		"phoneNumber": "+31 123456789",
		"password":    "SebWilder9999",
	}

	for _, customer := range customers {
		seedCustomer(
			customer["fname"],
			customer["lname"],
			customer["email"],
			customer["phoneNumber"],
			customer["password"],
			s)
	}

}

func seedAccount(id int, store storage.Storage) *model.Account {
	acc := model.NewAccount(id)

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s storage.Storage) {
	for i := 1; i <= 3; i++ {
		seedAccount(i, s)
	}
}
