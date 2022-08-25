package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kironono/pinkie/config"
	"github.com/kironono/pinkie/handler"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	db, cleanup, err := store.NewDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{"status": "ok"}`))
	})

	// jobs
	repo := registry.NewRepository(db)
	jobHandler := handler.NewJob(repo)
	mux.Get("/jobs/show", jobHandler.Show)

	return mux, cleanup, nil
}
