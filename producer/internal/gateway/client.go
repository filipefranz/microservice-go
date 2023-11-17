package gateway

import "github.com/filipefranz/microservice-go/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
