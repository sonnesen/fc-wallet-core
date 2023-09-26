package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Customer  *Customer
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(customer *Customer) (*Account, error) {
	account := &Account{
		ID:        uuid.New().String(),
		Customer:  customer,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Validate() error {
	if a.Customer == nil {
		return errors.New("customer is required")
	}
	return nil
}

func (a *Account) Credit(amount float64) error {
	err := a.ValidateAmount(amount)
	if err != nil {
		return err
	}
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Debit(amount float64) error {
	err := a.ValidateAmount(amount)
	if err != nil {
		return err
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return nil
}

// validate amount value
func (a *Account) ValidateAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}
