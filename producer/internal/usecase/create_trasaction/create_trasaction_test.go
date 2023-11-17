package createtrasaction

import (
	"testing"

	"github.com/filipefranz/microservice-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (t *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (a *AccountGatewayMock) Save(account *entity.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "johndoe@me.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("John Doe 2", "johndoe2@me.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccountGateway := &AccountGatewayMock{}
	mockAccountGateway.On("FindById", account1.ID).Return(account1, nil)
	mockAccountGateway.On("FindById", account2.ID).Return(account2, nil)

	mockTransactionGateway := &TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway)
	outPut, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outPut)
	mockAccountGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "FindById", 2)
	mockTransactionGateway.AssertExpectations(t)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)
}
