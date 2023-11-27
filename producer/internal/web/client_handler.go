package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/filipefranz/microservice-go/internal/usecase/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase create_client.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {

	var input create_client.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Erro 1", err)
		return
	}

	output, err := h.CreateClientUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro 2", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro 3", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
