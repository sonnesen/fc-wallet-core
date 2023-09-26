package account

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	CustomerID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	accountGateway  gateway.AccountGateway
	customerGateway gateway.CustomerGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, customerGateway gateway.CustomerGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountGateway:  accountGateway,
		customerGateway: customerGateway,
	}
}

func (u *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	customer, err := u.customerGateway.Get(input.CustomerID)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(customer)
	if err != nil {
		return nil, err
	}

	err = u.accountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
