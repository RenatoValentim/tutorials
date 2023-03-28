package entity

import "testing"

func TestNewClient(t *testing.T) {
	client, err := NewClient("fake_name", "fake_email")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if client.ID == "" {
		t.Errorf("unexpected empty ID")
	}

	if client.Name != "fake_name" {
		t.Errorf("unexpected name: %v", client.Name)
	}

	if client.Email != "fake_email" {
		t.Errorf("unexpected email: %v", client.Email)
	}

	if client.Points != 0 {
		t.Errorf("unexpected points: %v", client.Points)
	}
}
