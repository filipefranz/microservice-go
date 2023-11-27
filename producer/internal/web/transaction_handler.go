package web

import (
	"encoding/json"
	"net/http"

	"github.com/filipefranz/microservice-go/internal/usecase/create_transaction"
)

type WebTransactionCreated struct {
	CreateTransactionUseCase create_transaction.CreateTransactionUseCase
}

func NewWebTransactionCreated(createTransactionUseCase create_transaction.CreateTransactionUseCase) *WebTransactionCreated {
	return &WebTransactionCreated{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionCreated) TransactionCreated(w http.ResponseWriter, r *http.Request) {

	var input create_transaction.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	output, err := h.CreateTransactionUseCase.Execute(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
