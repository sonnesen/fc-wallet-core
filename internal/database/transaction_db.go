package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (t *TransactionDB) Save(transaction *entity.Transaction) error {
	stmt, err := t.DB.Prepare(`
		INSERT INTO transactions (id, account_from_id, account_to_id, amount, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.AccountFrom.ID,
		transaction.AccountTo.ID,
		transaction.Amount,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
