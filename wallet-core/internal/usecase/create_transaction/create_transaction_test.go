package create_transaction

import (
	"context"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/event"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/mocks"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	customer1, _ := entity.NewCustomer("John Doe", "j@j.com")
	account1 := entity.NewAccount(customer1)
	account1.Credit(100)

	customer2, _ := entity.NewCustomer("Jane Doe", "j@a.com")
	account2 := entity.NewAccount(customer2)
	account2.Credit(100)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	input := CreateTransactionInput{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        50,
	}

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)

	output, err := uc.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
