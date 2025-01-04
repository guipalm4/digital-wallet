package create_customer

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type CustomerGatewayMock struct {
	mock.Mock
}

func (m *CustomerGatewayMock) Get(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

func (m *CustomerGatewayMock) Save(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func TestCreteCustomerUseCase_Execute(t *testing.T) {
	customerGateway := &CustomerGatewayMock{}
	customerGateway.On("Save", mock.Anything).Return(nil)

	createCustomerUseCase := NewCreateCustomerUseCase(customerGateway)

	output, err := createCustomerUseCase.Execute(CreateCustomerInput{
		Name:  "John Doe",
		Email: "j@j.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j.com", output.Email)
	customerGateway.AssertExpectations(t)
	customerGateway.AssertNumberOfCalls(t, "Save", 1)
}
