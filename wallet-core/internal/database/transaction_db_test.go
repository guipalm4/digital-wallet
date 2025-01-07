package database

import (
	"database/sql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionDBSuite struct {
	suite.Suite
	db            *sql.DB
	customer      *entity.Customer
	customer2     *entity.Customer
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance decimal(12, 2), created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount decimal(12, 2), created_at date)")
	s.transactionDB = NewTransactionDB(db)
	customer, err := entity.NewCustomer("John Doe", "j@j.com")
	s.Nil(err)
	s.customer = customer
	customer2, err := entity.NewCustomer("Jane Doe", "j@a.com")
	s.Nil(err)
	s.customer2 = customer2
	accountFrom := entity.NewAccount(s.customer)
	accountFrom.Balance = 100
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(s.customer2)
	accountTo.Balance = 100
	s.accountTo = accountTo
	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBSuite))
}

func (s *TransactionDBSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 50)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
