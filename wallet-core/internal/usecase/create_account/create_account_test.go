package create_account

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

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j.com")
	customerMock := &CustomerGatewayMock{}
	customerMock.On("Get", customer.ID).Return(customer, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, customerMock)
	input := CreateAccountInput{
		CustomerID: customer.ID,
	}
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	customerMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	customerMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
