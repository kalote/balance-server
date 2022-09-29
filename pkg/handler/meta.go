package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) LiveProbe(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("liveness probe check")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProbeResponse{
		Status: "UP",
		Type:   "Live",
	})
}

func (h *Handler) ReadyProbe(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Readiness probe check")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProbeResponse{
		Status: "UP",
		Type:   "Ready",
	})
}
