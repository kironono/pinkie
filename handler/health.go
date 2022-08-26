package handler

import "net/http"

type HealthHandler interface {
	Show(http.ResponseWriter, *http.Request)
}

type health struct{}

func NewHealth() HealthHandler {
	return &health{}
}

func (h *health) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b := struct {
		Status string `json:"status"`
	}{
		"ok",
	}
	RespondJSON(ctx, w, b, http.StatusOK)
}
