package transaction

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

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Save(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	// arrange
	customer1, _ := entity.NewCustomer("John Doe", "john.doe@mail.com")
	customer2, _ := entity.NewCustomer("Jane Doe", "jane.doe@mail.com")

	account1, _ := entity.NewAccount(customer1)
	account1.Credit(1000)

	account2, _ := entity.NewAccount(customer2)
	account2.Credit(1000)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Get", account1.ID).Return(account1, nil)
	accountGatewayMock.On("Get", account2.ID).Return(account2, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Save", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDfrom: account1.ID,
		AccountIDto:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)

	// act
	output, err := uc.Execute(inputDto)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, output)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)

	accountGatewayMock.AssertNumberOfCalls(t, "Get", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
