package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/database"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/event"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/event/handler"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_account"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_customer"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_transaction"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/web"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/web/webserver"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/kafka"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/uow"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = runMigrations(dbUrl)

	if err != nil {
		panic(fmt.Errorf("error executing migrations: %w", err))
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	customerDb := database.NewCustomerDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createCustomerUseCase := create_customer.NewCreateCustomerUseCase(customerDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, customerDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(
		uow,
		eventDispatcher,
		transactionCreatedEvent,
		balanceUpdatedEvent)

	server := webserver.NewWebServer(":8080")

	customerHandler := web.NewWebCustomerHandler(*createCustomerUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	server.AddHandler("/customers", customerHandler.CreateCustomer)
	server.AddHandler("/accounts", accountHandler.CreateAccount)
	server.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running 🚀")
	server.Start()
}

func runMigrations(dbUrl string) error {

	migrationDSN := "mysql://" + dbUrl
	migrationPath := "file:///app/migrations"

	m, err := migrate.New(migrationPath, migrationDSN)
	if err != nil {
		panic(fmt.Errorf("Error on create migrator: %w", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("erro on apply migrations: %w", err))
	} else if err == migrate.ErrNoChange {
		log.Println("No pending migrations to be applied.")
	} else {
		log.Println("Migrations successfully applied!")
	}

	if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
		log.Printf("Errors when closing migrator (sourceErr: %v, dbErr: %v)", srcErr, dbErr)
	}
	return nil
}
