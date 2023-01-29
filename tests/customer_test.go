package tests

import (
	"fmt"
	"testing"

	"github.com/DanielMolinaR/bank-project/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	customer, err := model.NewCustomer("test_name", "test_last_name", "test@mail.com", "+34 123456789", "test_password")
	assert.Nil(t, err)

	fmt.Printf("%v\n", customer)

}
