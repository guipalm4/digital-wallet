package database

import (
	"database/sql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{DB: db}
}

func (c *ClientDB) Get(id string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	stmt, err := c.DB.Prepare("SELECT id, name, created_at FROM customers WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	if err := row.Scan(&customer.ID, &customer.Name, &customer.CreatedAt); err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *ClientDB) Save(customer *entity.Customer) error {
	stmt, err := c.DB.Prepare("INSERT INTO customers (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
