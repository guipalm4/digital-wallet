package gateway

import "github.com/guipalm4/digital-wallet/wallet-core/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
