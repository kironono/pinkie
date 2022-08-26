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
	repo := registry.NewRepository(db)

	// health
	mux.Route("/health", func(r chi.Router) {
		h := handler.NewHealth()

		r.Get("/", h.Show)
	})

	// users
	mux.Route("/users", func(r chi.Router) {
		h := handler.NewUser(repo)

		r.Get("/", h.List)
		r.Get("/{id:\\d+}", h.Show)
	})

	// jobs
	mux.Route("/jobs", func(r chi.Router) {
		h := handler.NewJob(repo)

		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{id:\\d+}", func(r chi.Router) {
			r.Get("/", h.Show)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})

	// metric
	mux.Route("/v1/metric/{jobSlug}", func(r chi.Router) {
		h := handler.NewMetric(repo)
		r.Post("/", h.Create)
	})

	return mux, cleanup, nil
}
