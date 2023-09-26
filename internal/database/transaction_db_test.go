package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	customer1     *entity.Customer
	customer2     *entity.Customer
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at datetime, updated_at datetime)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance float, created_at datetime, updated_at datetime)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount float, created_at datetime, updated_at datetime)")
	s.transactionDB = NewTransactionDB(db)
	s.customer1, _ = entity.NewCustomer("John Doe", "john.doe@mail.com")
	s.customer2, _ = entity.NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountFrom, _ := entity.NewAccount(s.customer1)
	accountTo, _ := entity.NewAccount(s.customer2)
	accountFrom.Balance = 1000
	accountTo.Balance = 1000
	s.accountFrom = accountFrom
	s.accountTo = accountTo
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
	s.db.Close()
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestSave() {
	transaction, _ := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	err := s.transactionDB.Save(transaction)
	s.Nil(err)
}
