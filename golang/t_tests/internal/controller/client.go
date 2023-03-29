package controller

import (
	"encoding/json"
	"net/http"
	"t_tests/internal/database"
	"t_tests/internal/usecase"
)

func (b *BaseHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var content usecase.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	repo := database.NewClientRepository(b.db)
	uc := usecase.NewCreateClientUsecase(repo)
	_, err = uc.Execute(&content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
