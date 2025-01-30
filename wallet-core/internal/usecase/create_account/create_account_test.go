package create_account

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	customer, _ := entity.NewCustomer("John Doe", "j@j")
	customerMock := &mocks.CustomerGatewayMock{}
	customerMock.On("Get", customer.ID).Return(customer, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, customerMock)
	inputDto := CreateAccountInput{
		CustomerID: customer.ID,
	}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	// asssert valid uuid
	customerMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	customerMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
