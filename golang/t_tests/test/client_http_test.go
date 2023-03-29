package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"t_tests/internal/controller"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	sqlStatemants := `
	CREATE TABLE IF NOT EXISTS clients (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		pointers INTEGER NOT NULL
	);
	`
	_, err = db.Exec(sqlStatemants)
	if err != nil {
		panic(err)
	}

	return db
}

func tearDown(db *sql.DB) {
	db.Exec(`DROP TABLE clients;`)
	db.Close()
}

func TestCreateClientHandler(t *testing.T) {
	db := setupDb()
	defer tearDown(db)
	controller := controller.NewBaseHandler(db)

	t.Run(`Should create a client`, func(t *testing.T) {
		data := `{"name": "fake_name", "email": "fake_email"}`
		reader := strings.NewReader(data)
		request, _ := http.NewRequest("POST", "/clients", reader)
		response := httptest.NewRecorder()
		controller.CreateClientHandler(response, request)
		if response.Code != http.StatusCreated {
			t.Errorf("expected status code %d got %d", http.StatusCreated, response.Code)
		}
	})
}
