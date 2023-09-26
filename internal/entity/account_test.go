package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")

	// act
	account, err := NewAccount(customer)

	// assert
	assert.NotNil(t, account)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, account.Customer.ID)
}

func TestCreditAccount(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account, _ := NewAccount(customer)

	// act
	account.Credit(100)

	// assert
	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account, _ := NewAccount(customer)
	account.Credit(100)

	// act
	account.Debit(50)

	// assert
	assert.Equal(t, 50.0, account.Balance)
}
