package entity

import (
	"github.com/google/uuid"
)

type Client struct {
	ID     string
	Name   string
	Email  string
	Points int
}

func NewClient(name, email string) (*Client, error) {
	return &Client{
		ID:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Points: 0,
	}, nil
}
