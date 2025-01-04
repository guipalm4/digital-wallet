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
