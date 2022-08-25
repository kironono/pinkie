package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kironono/pinkie/entity"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/usecase"
)

type JobHandler interface {
	Show(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
}

type job struct {
	uc usecase.Job
}

func NewJob(repo registry.Repository) JobHandler {
	uc := usecase.NewJob(repo.NewJob())
	return &job{
		uc: uc,
	}
}

func (j *job) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	job, err := j.uc.Show(ctx, entity.JobID(id))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, job, http.StatusOK)
}

func (j *job) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jobs, err := j.uc.List(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, jobs, http.StatusOK)
}
