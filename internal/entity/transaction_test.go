package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test creation transaction
func TestCreateTransaction(t *testing.T) {
	// arrange
	customerFrom, _ := NewCustomer("John Doe", "john.doe@mail.com")
	accountFrom, _ := NewAccount(customerFrom)
	accountFrom.Credit(200)

	customerTo, _ := NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountTo, _ := NewAccount(customerTo)

	// act
	transaction, err := NewTransaction(accountFrom, accountTo, 100)

	// assert
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
	assert.Equal(t, accountFrom.ID, transaction.AccountFrom.ID)
	assert.Equal(t, accountTo.ID, transaction.AccountTo.ID)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, 100.0, accountFrom.Balance)
	assert.Equal(t, 100.0, accountTo.Balance)
}

// test creation transaction when account from has insufficient funds
func TestCreateTransactionWhenAccountFromHasInsufficientFunds(t *testing.T) {
	// arrange
	customerFrom, _ := NewCustomer("John Doe", "john.doe@mail.com")
	accountFrom, _ := NewAccount(customerFrom)

	customerTo, _ := NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountTo, _ := NewAccount(customerTo)

	// act
	transaction, err := NewTransaction(accountFrom, accountTo, 100)

	// assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "account from has insufficient funds", err.Error())
}

// test creation transaction when account from is nil
func TestCreateTransactionWhenAccountFromIsNil(t *testing.T) {
	// arrange
	customerTo, _ := NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountTo, _ := NewAccount(customerTo)

	// act
	transaction, err := NewTransaction(nil, accountTo, 100)

	// assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "account from is required", err.Error())
}

// test creation transaction when account to is nil
func TestCreateTransactionWhenAccountToIsNil(t *testing.T) {
	// arrange
	customerFrom, _ := NewCustomer("John Doe", "john.doe@mail.com")
	accountFrom, _ := NewAccount(customerFrom)

	// act
	transaction, err := NewTransaction(accountFrom, nil, 100)

	// assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "account to is required", err.Error())
}

// test creation transaction when amount is zero
func TestCreateTransactionWhenAmountIsZero(t *testing.T) {
	// arrange
	customerFrom, _ := NewCustomer("John Doe", "john.doe@mail.com")
	accountFrom, _ := NewAccount(customerFrom)

	customerTo, _ := NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountTo, _ := NewAccount(customerTo)

	// act
	transaction, err := NewTransaction(accountFrom, accountTo, 0)

	// assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}

// test creation transaction when amount is less than zero
func TestCreateTransactionWhenAmountIsLessThanZero(t *testing.T) {
	// arrange
	customerFrom, _ := NewCustomer("John Doe", "john.doe@mail.com")
	accountFrom, _ := NewAccount(customerFrom)

	customerTo, _ := NewCustomer("Jane Doe", "jane.doe@mail.com")
	accountTo, _ := NewAccount(customerTo)

	// act
	transaction, err := NewTransaction(accountFrom, accountTo, -100)

	// assert
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}
