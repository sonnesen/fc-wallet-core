package account

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

type CustomerGatewayMock struct {
	mock.Mock
}

func (m *CustomerGatewayMock) Save(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *CustomerGatewayMock) Get(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

func TestCreateAccountUsesCase_Execute(t *testing.T) {
	// arrange
	customer, _ := entity.NewCustomer("John Doe", "john.doe@mail.com")

	customerGatewayMock := &CustomerGatewayMock{}
	customerGatewayMock.On("Get", customer.ID).Return(customer, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, customerGatewayMock)
	inputDto := CreateAccountInputDTO{
		CustomerID: customer.ID,
	}

	// act
	output, err := uc.Execute(inputDto)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, output)
	customerGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	customerGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
