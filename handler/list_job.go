package handler

import (
	"context"
	"net/http"

	"github.com/kironono/pinkie/entity"
)

type ListJobsService interface {
	ListJobs(ctx context.Context) (entity.Jobs, error)
}

type ListJobs struct {
	Service ListJobsService
}

type job struct {
	ID   entity.JobID `json:"id"`
	Name string       `json:"name"`
	Slug string       `json:"slug"`
}

func (l *ListJobs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jobs, err := l.Service.ListJobs(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := []job{}
	for _, j := range jobs {
		rsp = append(rsp, job{
			ID: j.ID,
		})
	}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
