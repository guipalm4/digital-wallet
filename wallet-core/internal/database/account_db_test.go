package database

import (
	"database/sql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	customer  *entity.Customer
}

func (s *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance decimal(12, 2), created_at date, updated_at date)")
	s.accountDB = NewAccountDB(db)
	s.customer, _ = entity.NewCustomer("John Doe", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.db.Exec("INSERT INTO customers (id, name, email, created_at) VALUES (?, ?, ?, ?)",
		s.customer.ID, s.customer.Name, s.customer.Email, s.customer.CreatedAt,
	)

	account := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Customer.ID, accountDB.Customer.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Customer.Name, accountDB.Customer.Name)
	s.Equal(account.Customer.Email, accountDB.Customer.Email)
}
