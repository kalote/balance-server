package handler

import (
	"encoding/json"
	"net/http"
)

func (h *handler) sendProbe(probeType string, w http.ResponseWriter) {
	h.logger.Info("liveness probe check")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProbeResponse{
		Status: "OK",
		Type:   probeType,
	})
}
