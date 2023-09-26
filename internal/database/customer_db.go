package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type CustomerDB struct {
	DB *sql.DB
}

func NewCustomerDB(db *sql.DB) *CustomerDB {
	return &CustomerDB{
		DB: db,
	}
}

func (c *CustomerDB) Get(id string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at, updated_at FROM customers WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerDB) Save(customer *entity.Customer) error {
	stmt, err := c.DB.Prepare("INSERT INTO customers (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
