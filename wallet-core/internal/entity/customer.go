package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name, email string) (*Customer, error) {
	customer := &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := customer.Validate()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.New("'name' is required")
	}
	if c.Email == "" {
		return errors.New("'email' is required")
	}
	return nil
}

func (c *Customer) Update(name, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	return c.Validate()
}

func (c *Customer) AddAccount(account *Account) error {

	if account.Customer.ID != c.ID {
		return errors.New("account does not belong to this customer")
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
