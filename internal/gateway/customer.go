package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type CustomerGateway interface {
	Save(customer *entity.Customer) error
	Get(id string) (*entity.Customer, error)
}
