package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalote/balance-server/pkg/internal/validation"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := validation.CheckEthAddress(vars); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["ethAddress"]))
}
