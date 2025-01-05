package gateway

import "github.com/guipalm4/digital-wallet/wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
