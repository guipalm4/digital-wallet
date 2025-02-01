package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account := NewAccount(customer)
	assert.NotNil(t, account)
	assert.Equal(t, customer.ID, account.Customer.ID)
}

func TestCreateAccountWithNilCustomer(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account := NewAccount(customer)
	account.Credit(100)
	assert.Equal(t, 1100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account := NewAccount(customer)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, 1050.0, account.Balance)
}
