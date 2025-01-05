package create_account

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/gateway"
)

type CreateAccountInput struct {
	CustomerID string
}

type CreateAccountOutput struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway  gateway.AccountGateway
	CustomerGateway gateway.CustomerGateway
}

func NewCreateAccountUseCase(
	accountGateway gateway.AccountGateway,
	customerGateway gateway.CustomerGateway) *CreateAccountUseCase {

	return &CreateAccountUseCase{
		AccountGateway:  accountGateway,
		CustomerGateway: customerGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	customer, err := uc.CustomerGateway.Get(input.CustomerID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(customer)

	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{
		ID: account.ID,
	}, nil
}
