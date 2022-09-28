package handler

import (
	"net/http"
)

type ProbeResponse struct {
	Status string `json:"status"`
	Type   string `json:"type"`
}

func (h *handler) LiveProbe(w http.ResponseWriter, r *http.Request) {
	h.sendProbe("live", w)
}

func (h *handler) ReadyProbe(w http.ResponseWriter, r *http.Request) {
	h.sendProbe("ready", w)
}
