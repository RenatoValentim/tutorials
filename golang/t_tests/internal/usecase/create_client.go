package usecase

import "t_tests/internal/entity"

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID     string
	Name   string
	Email  string
	Points int
}

type CreateClientUsecase struct {
	ClientRepository entity.ClientRepositoryInterface
}

func NewCreateClientUsecase(clientRepository entity.ClientRepositoryInterface) *CreateClientUsecase {
	return &CreateClientUsecase{
		ClientRepository: clientRepository,
	}
}

func (c *CreateClientUsecase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != err {
		return nil, err
	}

	err = c.ClientRepository.Save(client)
	if err != err {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:     client.ID,
		Name:   client.Name,
		Email:  client.Email,
		Points: client.Points,
	}, nil
}
