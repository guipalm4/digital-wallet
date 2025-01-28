package create_transaction

import (
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/event"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	customer1, _ := entity.NewCustomer("John Doe", "j@j.com")
	account1 := entity.NewAccount(customer1)
	account1.Credit(100)

	customer2, _ := entity.NewCustomer("Jane Doe", "j@a.com")
	account2 := entity.NewAccount(customer2)
	account2.Credit(100)

	mockAccountGateway := &AccountGatewayMock{}
	mockAccountGateway.On("Get", account1.ID).Return(account1, nil)
	mockAccountGateway.On("Get", account2.ID).Return(account2, nil)

	mockTransactionGateway := &TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInput{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        50,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway, dispatcher, event)

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	mockAccountGateway.AssertExpectations(t)
	mockTransactionGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "Get", 2)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)
}
