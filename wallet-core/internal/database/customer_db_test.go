package database

import (
	"database/sql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *CustomerDB
}

func (s *ClientDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
	db.Exec("Create table customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewCustomerDB(db)
}

func (s *ClientDBTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	customer := &entity.Customer{
		ID:    "1",
		Name:  "John Doe",
		Email: "j@j.com",
	}
	err := s.clientDB.Save(customer)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("John Doe", "j@j.com")
	s.clientDB.Save(customer)

	clientDB, err := s.clientDB.Get(customer.ID)
	s.Nil(err)
	s.Equal(customer.ID, clientDB.ID)
	s.Equal(customer.Name, clientDB.Name)
	s.Equal(customer.Email, clientDB.Email)
}
