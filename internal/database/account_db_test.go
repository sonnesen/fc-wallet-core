package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	customer  *entity.Customer
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance float, created_at datetime, updated_at datetime)")
	s.accountDB = NewAccountDB(db)
	s.customer, _ = entity.NewCustomer("John Doe", "john.doe@mail.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
	s.db.Close()
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, _ := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestGet() {
	s.db.Exec("INSERT INTO customers (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		s.customer.ID, s.customer.Name, s.customer.Email, s.customer.CreatedAt, s.customer.UpdatedAt)

	account, _ := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Customer.ID, accountDB.Customer.ID)
	s.Equal(account.Customer.Name, accountDB.Customer.Name)
	s.Equal(account.Customer.Email, accountDB.Customer.Email)
	s.Equal(account.Balance, accountDB.Balance)
}
