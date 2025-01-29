package web

import (
	"encoding/json"
	"github.com/guipalm4/digital-wallet/wallet-core/internal/usecase/create_customer"
	"net/http"
)

type WebCustomerHandler struct {
	CreateCustomerUseCase create_customer.CreateCustomerUseCase
}

func NewWebCustomerHandler(createCustomerUseCase create_customer.CreateCustomerUseCase) *WebCustomerHandler {
	return &WebCustomerHandler{
		CreateCustomerUseCase: createCustomerUseCase,
	}
}

func (h *WebCustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var input create_customer.CreateCustomerInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateCustomerUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
