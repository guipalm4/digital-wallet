package gateway

import "github.com/guipalm4/digital-wallet/wallet-core/internal/entity"

type CustomerGateway interface {
	Get(id string) (*entity.Customer, error)
	Save(customer *entity.Customer) error
}
