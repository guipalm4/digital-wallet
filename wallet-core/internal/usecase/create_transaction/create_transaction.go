package create_transaction

import (
	"context"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/entity"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/gateway"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/uow"
	"log"
)

type CreateTransactionInput struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutput struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutput struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.IEventDispatcher
	TransactionCreated events.IEvent
	BalanceUpdated     events.IEvent
}

func NewCreateTransactionUseCase(
	Uow uow.UowInterface,
	eventDispatcher events.IEventDispatcher,
	transactionCreated events.IEvent,
	balanceUpdated events.IEvent,
) *CreateTransactionUseCase {

	return &CreateTransactionUseCase{
		Uow:                Uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
		BalanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInput) (*CreateTransactionOutput, error) {

	output := &CreateTransactionOutput{}
	balanceUpdatedOutput := &BalanceUpdatedOutput{}

	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)
		log.Printf("Recovering account ID from: %s", input.AccountIDFrom)
		accountFrom, err := accountRepository.Get(input.AccountIDFrom)
		if err != nil {
			return err
		}
		log.Printf("Recovering account ID to %s", input.AccountIDTo)
		accountTo, err := accountRepository.Get(input.AccountIDTo)
		if err != nil {
			return err
		}
		log.Printf("Creating transaction. Account from: %s, Account to: %s, Amount: %f", input.AccountIDFrom, input.AccountIDTo, input.Amount)
		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}
		log.Printf("Updating balance account from: %s", input.AccountIDFrom)
		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}
		log.Printf("Updating balance account to: %s", input.AccountIDTo)
		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		log.Printf("Creating transaction on database")
		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
		balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
		balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance

		return nil
	})

	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	log.Printf("Dispatching event TransactionCreated with payload: %v", output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	log.Printf("Dispatching event BalanceUpdated with payload: %v", balanceUpdatedOutput)
	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)
	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
