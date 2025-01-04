package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNemCustomer(t *testing.T) {
	customer, err := NewCustomer("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "j@j.com", customer.Email)
}

func TestCreateNemCustomerWhenArgsAreInvalid(t *testing.T) {
	customer, err := NewCustomer("", "")
	assert.NotNil(t, err)
	assert.Nil(t, customer)
}

func TestUpdateCustomer(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "j@j.com")
	err := customer.Update("Jane Doe", "j@g.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe", customer.Name)
	assert.Equal(t, "j@g.com", customer.Email)
}

func TestUpdateWhenArgsAreInvalid(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "j@j.com")
	err := customer.Update("", "j@g.com")
	assert.Error(t, err, "'name' is required")
}

func TestAddAccountToCustomer(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "j@j.com")
	account := NewAccount(customer)
	err := customer.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(customer.Accounts))
}
