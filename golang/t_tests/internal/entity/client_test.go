package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("fake_name", "fake_email")

	assert.Nil(t, err)
	assert.NotEqual(t, "", client.ID)
	assert.Equal(t, "fake_name", client.Name)
	assert.Equal(t, "fake_email", client.Email)
	assert.Equal(t, 0, client.Points)
}

func TestNewClientWithInvalidName(t *testing.T) {
	client, err := NewClient("", "fake_email")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client name is required")
}

func TestNewClientWithInvalidEmail(t *testing.T) {
	client, err := NewClient("fake_name", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "client email is required")
}

func TestAddPoints(t *testing.T) {
	client, err := NewClient("fake_name", "fake_email")

	client.AddPoints(10)

	assert.Nil(t, err)
	assert.NotEqual(t, "", client.ID)
	assert.Equal(t, "fake_name", client.Name)
	assert.Equal(t, "fake_email", client.Email)
	assert.Equal(t, 10, client.Points)
}
