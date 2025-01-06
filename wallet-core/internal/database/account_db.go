package database

import (
	"database/sql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (a *AccountDB) Get(id string) (*entity.Account, error) {
	var account entity.Account
	var customer entity.Customer

	account.Customer = &customer

	stmt, err := a.DB.Prepare(
		"SELECT a.id, " +
			"a.customer_id, " +
			"a.balance, " +
			"a.created_at, " +
			"c.id, c.name, " +
			"c.email, " +
			"c.created_at " +
			"FROM accounts a " +
			"Inner Join customers c " +
			"ON a.customer_id = c.id " +
			"WHERE a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.Customer.ID,
		&account.Balance,
		&account.CreatedAt,
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare(
		"INSERT INTO accounts (id, customer_id, balance, created_at) " +
			"VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		account.ID,
		account.Customer.ID,
		account.Balance,
		account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
