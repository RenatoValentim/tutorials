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

func TestAddPointsMinorOrEqualsToZero(t *testing.T) {
	client, _ := NewClient("fake_name", "fake_email")

	err := client.AddPoints(0)

	assert.NotNil(t, err)
	assert.Error(t, err, "points must be greater than zero")
	assert.NotEqual(t, "", client.ID)
	assert.Equal(t, "fake_name", client.Name)
	assert.Equal(t, "fake_email", client.Email)
	assert.Equal(t, 0, client.Points)
}

func FuzzAddPointsBatch(f *testing.F) {
	points := []int{2, 4, 6, 8, 10}
	for _, point := range points {
		f.Add(point)
	}

	f.Fuzz(func(t *testing.T, points int) {
		client, _ := NewClient("fake_name", "fake_email")

		err := client.AddPoints(points)

		if err != nil {
			return
		}

		if client.Points != points {
			t.Errorf("Points expected: %d, got: %d", points, client.Points)
		}
	})
}
