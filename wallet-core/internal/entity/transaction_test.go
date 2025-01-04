package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	customer1, _ := NewCustomer("John Doe", "j@j.com")
	account1 := NewAccount(customer1)
	customer2, _ := NewCustomer("Jane Doe", "j@a.com")
	account2 := NewAccount(customer2)

	account1.Credit(100)
	account2.Credit(100)

	transaction, err := NewTransaction(account1, account2, 50)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 50.0, account1.Balance)
	assert.Equal(t, 150.0, account2.Balance)
}

func TestNewTransactionWhenInsufficientBalance(t *testing.T) {
	customer1, _ := NewCustomer("John Doe", "j@j.com")
	account1 := NewAccount(customer1)
	customer2, _ := NewCustomer("Jane Doe", "j@a.com")
	account2 := NewAccount(customer2)

	account1.Credit(100)
	account2.Credit(100)

	transaction, err := NewTransaction(account1, account2, 200)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 100.0, account1.Balance)
	assert.Equal(t, 100.0, account2.Balance)
}

func TestNewTransactionWhenAmountIsZero(t *testing.T) {
	customer1, _ := NewCustomer("John Doe", "j@j.com")
	account1 := NewAccount(customer1)
	customer2, _ := NewCustomer("Jane Doe", "j@a.com")
	account2 := NewAccount(customer2)

	account1.Credit(100)
	account2.Credit(100)

	transaction, err := NewTransaction(account1, account2, 0)
	assert.NotNil(t, err)
	assert.Error(t, err, "'amount' must be greater than 0")
	assert.Nil(t, transaction)
	assert.Equal(t, 100.0, account1.Balance)
	assert.Equal(t, 100.0, account2.Balance)
}
