package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//eventDispatcher := events.NewEventDispatcher()
	//transactionCreatedEvent := event.NewTransactionCreated()
	////eventDispatcher.Register("TransactionCreated", handler)
	//
	//customerDb := database.NewCustomerDB(db)
	//accountDb := database.NewAccountDB(db)
	//transactionDb := database.NewTransactionDB(db)
	//
	//createCustomerUseCase := create_customer.NewCreateCustomerUseCase(customerDb)
	//createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, customerDb)
	//createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

}
