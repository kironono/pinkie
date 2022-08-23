package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

func NewMux(ctx context.Context) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{"status": "ok"}`))
	})

	return mux, nil
}
