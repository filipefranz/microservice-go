package mocks

import (
	"context"

	"github.com/filipefranz/microservice-go/pkg/uow"
	"github.com/stretchr/testify/mock"
)

type UowMock struct {
	mock.Mock
}

func (u *UowMock) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	args := u.Called(ctx, fn)
	return args.Error(0)
}

func (u *UowMock) CommitOrRollback() error {
	args := u.Called()
	return args.Error(0)
}

func (u *UowMock) Rollback() error {
	args := u.Called()
	return args.Error(0)
}

func (u *UowMock) Register(name string, fc uow.RepositoryFactory) {
	u.Called(name, fc)
}

func (u *UowMock) UnRegister(name string) {
	u.Called(name)
}

func (u *UowMock) GetRepository(ctx context.Context, name string) (interface{}, error) {
	args := u.Called(ctx, name)
	return args.Get(0), args.Error(1)
}
