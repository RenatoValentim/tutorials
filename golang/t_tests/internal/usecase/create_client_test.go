package usecase

import (
	"t_tests/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientRepositoryMock struct {
	mock.Mock
}

func (c *ClientRepositoryMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

func TestCreateClientUsecase_Execute(t *testing.T) {
	mockRepo := new(ClientRepositoryMock)
	mockRepo.On("Save", mock.Anything).Return(nil)
	createClientUsecase := NewCreateClientUsecase(mockRepo)
	input := CreateClientInputDTO{
		Name:  "fake_name",
		Email: "fake_email",
	}
	output, err := createClientUsecase.Execute(&input)
	assert.Nil(t, err)
	assert.Equal(t, "fake_name", output.Name)
	assert.Equal(t, "fake_email", output.Email)
	assert.Equal(t, 0, output.Points)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "Save", 1)
}
