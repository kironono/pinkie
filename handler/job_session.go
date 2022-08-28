package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/registry"
	"github.com/kironono/pinkie/usecase"
)

type JobSessionHandler interface {
	List(http.ResponseWriter, *http.Request)
}

type jobSession struct {
	uc usecase.JobSession
}

func NewJobSession(repo registry.Repository) JobSessionHandler {
	uc := usecase.NewJobSession(repo.NewJobSession())
	return &jobSession{
		uc: uc,
	}
}

func (j *jobSession) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jobID, _ := strconv.Atoi(chi.URLParam(r, "jobID"))

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	per, err := strconv.Atoi(r.URL.Query().Get("per"))
	if err != nil || per < 1 {
		per = DEFAULT_PER_PAGE_NUM
	}
	order := "created_at desc"

	jobs, err := j.uc.List(ctx, model.JobID(jobID), model.PageNum(page), model.PerPageNum(per), model.Order(order))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, jobs, http.StatusOK)
}
