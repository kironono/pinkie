package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/usecase"
)

type MetricHandler interface {
	Create(http.ResponseWriter, *http.Request)
}

type metric struct {
	uc       usecase.Metric
	validate *validator.Validate
}

func NewMetric(repo registry.Repository) MetricHandler {
	return &metric{
		uc:       usecase.NewMetric(repo.NewMetric()),
		validate: validator.New(),
	}
}

func (m *metric) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jobSlug := chi.URLParam(r, "jobSlug")

	var b struct {
		Timestamp time.Time `json:"timestamp" validate:"required"`
		Status    string    `json:"status" validate:"required"`
		Task      string    `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := m.validate.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := m.uc.Create(ctx, jobSlug, b.Timestamp, b.Status, b.Task); err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			RespondJSON(ctx, w, &ErrResponse{
				Message: "Not Found",
			}, http.StatusNotFound)
		case errors.Is(err, model.ErrOpendJobSessionNotFound):
			RespondJSON(ctx, w, &ErrResponse{
				Message: "Bad Request",
			}, http.StatusBadRequest)
		default:
			log.Printf("%s\n", err)
			RespondJSON(ctx, w, &ErrResponse{
				Message: err.Error(),
			}, http.StatusInternalServerError)
		}
		return
	}
	RespondJSON(ctx, w, nil, http.StatusCreated)
}
