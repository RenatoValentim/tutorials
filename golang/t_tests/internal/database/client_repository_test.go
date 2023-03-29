package database

import (
	"database/sql"
	"t_tests/internal/entity"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestClientRepository_Save(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	defer db.Close()
	assert.Nil(t, err)

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS clients (
		id TEXT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		points INTEGER NOT NULL
	);
	`
	_, err = db.Exec(sqlStatement)
	assert.Nil(t, err)

	clientRepository := NewClientRepository(db)
	client := &entity.Client{
		ID:     "fake_id",
		Name:   "fake_name",
		Email:  "fake_email",
		Points: 0,
	}
	err = clientRepository.Save(client)
	assert.Nil(t, err)

	var id, name, email string
	var points int
	err = db.QueryRow(`
		SELECT id, name, email, points FROM clients
		WHERE id = $1;
		`,
		client.ID,
	).Scan(
		&id,
		&name,
		&email,
		&points,
	)
	assert.Nil(t, err)
	assert.Equal(t, id, client.ID)
	assert.Equal(t, name, client.Name)
	assert.Equal(t, email, client.Email)
	assert.Equal(t, points, client.Points)
}
