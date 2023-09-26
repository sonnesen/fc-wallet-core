package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// create a new transaction
func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	err = transaction.Process()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// validate transaction
func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("account from is required")
	}
	if t.AccountTo == nil {
		return errors.New("account to is required")
	}
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("account from has insufficient funds")
	}
	return nil
}

// process transaction
func (t *Transaction) Process() error {
	err := t.AccountFrom.Debit(t.Amount)
	if err != nil {
		return err
	}
	err = t.AccountTo.Credit(t.Amount)
	if err != nil {
		return err
	}
	return nil
}
