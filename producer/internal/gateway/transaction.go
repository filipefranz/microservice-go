package gateway

import "github.com/filipefranz/microservice-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
