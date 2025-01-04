package entity

import "time"

type Account struct {
	ID        string
	Customer  *Customer
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(customer *Customer) *Account {

	if customer == nil {
		return nil
	}

	return &Account{
		ID:        customer.ID,
		Customer:  customer,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}
