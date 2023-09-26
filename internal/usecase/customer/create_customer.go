package usecase_customer

import (
	"time"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type CustomerGateway interface {
	CreateCustomer(name, email string) (*CreateCustomerOutputDTO, error)
}

type CreateCustomerInputDTO struct {
	Name  string
	Email string
}

type CreateCustomerOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateCustomerUseCase struct {
	CustomerGateway gateway.CustomerGateway
}

func NewCreateCustomerUseCase(customerGateway gateway.CustomerGateway) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		CustomerGateway: customerGateway,
	}
}

func (u *CreateCustomerUseCase) Execute(input *CreateCustomerInputDTO) (*CreateCustomerOutputDTO, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = u.CustomerGateway.Save(customer)
	if err != nil {
		return nil, err
	}

	return &CreateCustomerOutputDTO{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}
