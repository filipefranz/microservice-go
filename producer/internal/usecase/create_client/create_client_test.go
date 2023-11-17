package createclient

import (
	"testing"

	"github.com/filipefranz/microservice-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (c *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (c *ClientGatewayMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(clientGatewayMock)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "John Doe",
		Email: "johndoe@me.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "johndoe@me.com", output.Email)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
