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
