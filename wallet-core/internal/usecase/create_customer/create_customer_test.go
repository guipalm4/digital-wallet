package create_customer

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreteCustomerUseCase_Execute(t *testing.T) {
	customerGateway := &mocks.CustomerGatewayMock{}
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
