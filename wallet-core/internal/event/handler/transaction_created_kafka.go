package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/guipalm4/digital-wallet/wallet-core/pkg/events"
	"github.com/guipalm4/digital-wallet/wallet-core/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.IEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "transactions")
	messageJson, err := json.MarshalIndent(message.GetPayload(), "", "    ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println("TransactionCreatedKafkaHandler: ", string(messageJson))
}
