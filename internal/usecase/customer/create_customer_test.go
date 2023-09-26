package usecase_customer

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	// arrange
	m := &CustomerGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateCustomerUseCase(m)

	output, err := uc.Execute(&CreateCustomerInputDTO{
		Name:  "John Doe",
		Email: "john.doe@mail.com",
	})

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "john.doe@mail.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
