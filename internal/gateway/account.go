package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	Get(id string) (*entity.Account, error)
}
