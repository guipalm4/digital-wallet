package create_customer

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/gateway"
	"time"
)

type CreateCustomerInput struct {
	Name  string
	Email string
}

type CreateCustomerOutput struct {
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

func (uc *CreateCustomerUseCase) Execute(input CreateCustomerInput) (*CreateCustomerOutput, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = uc.CustomerGateway.Save(customer)
	if err != nil {
		return nil, err
	}

	return &CreateCustomerOutput{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}
