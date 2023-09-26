package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CustomerDBTestSuite struct {
	suite.Suite
	db         *sql.DB
	customerDB *CustomerDB
}

func (s *CustomerDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	s.customerDB = NewCustomerDB(db)
}

func (s *CustomerDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE customers")
	s.db.Close()
}

func TestCustomerDBTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBTestSuite))
}

func (s *CustomerDBTestSuite) TestSave() {
	customer, _ := entity.NewCustomer("John Doe", "john.doe@mail.com")
	err := s.customerDB.Save(customer)
	s.Nil(err)
}

func (s *CustomerDBTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("John Doe", "john.doe@mail.com")
	s.customerDB.Save(customer)

	customerDB, err := s.customerDB.Get(customer.ID)

	s.Nil(err)
	s.Equal(customer.ID, customerDB.ID)
	s.Equal(customer.Name, customerDB.Name)
	s.Equal(customer.Email, customerDB.Email)
}
