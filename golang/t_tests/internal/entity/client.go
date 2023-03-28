package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Client struct {
	ID     string
	Name   string
	Email  string
	Points int
}

func NewClient(name, email string) (*Client, error) {
	if name == "" || email == "" {
		return nil, errors.New("client name and email is required")
	}

	return &Client{
		ID:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Points: 0,
	}, nil
}

func (c *Client) AddPoints(points int) error {
	if points <= 0 {
		return errors.New("points must be greater than zero")
	}

	c.Points = points

	return nil
}
