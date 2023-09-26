package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustomer(t *testing.T) {
	// arrange

	// act
	customer, err := NewCustomer("John Doe", "john.doe@mail.com")

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "john.doe@mail.com", customer.Email)
}

func TestCreateNewCustomerWhenArgsAreInvalid(t *testing.T) {
	// arrange

	// act
	customer, err := NewCustomer("", "")

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, customer)
}

func TestUpdateCustomer(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")

	// act
	err := customer.Update("John Doe Updated", "john.doe@email.com")

	// assert
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Updated", customer.Name)
	assert.Equal(t, "john.doe@email.com", customer.Email)
}

func TestUpdateCustomerWhenArgsAreInvalid(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")

	// act
	err := customer.Update("", "john.doe@email.com")

	// assert
	assert.Error(t, err, "name is required")
}

func TestAddAccountToCustomer(t *testing.T) {
	// arrange
	customer, _ := NewCustomer("John Doe", "john.doe@mail.com")
	account, _ := NewAccount(customer)

	// act
	err := customer.AddAccount(account)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(customer.Accounts))
}
