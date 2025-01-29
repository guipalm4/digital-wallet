package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/database"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/event"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_account"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_customer"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_transaction"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/web"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/web/webserver"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)
	customerDb := database.NewCustomerDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createCustomerUseCase := create_customer.NewCreateCustomerUseCase(customerDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, customerDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	server := webserver.NewWebServer(":3000")

	customerHandler := web.NewWebCustomerHandler(*createCustomerUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	server.AddHandler("/customers", customerHandler.CreateCustomer)
	server.AddHandler("/accounts", accountHandler.CreateAccount)
	server.AddHandler("/transactions", transactionHandler.CreateTransaction)

	server.Start()
}
