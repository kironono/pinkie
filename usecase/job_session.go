package usecase

import (
	"context"
	"fmt"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type JobSession interface {
	List(context.Context, model.JobID, model.PageNum, model.PerPageNum, model.Order) (model.JobSessions, error)
}

type jobSession struct {
	repo repository.JobSession
}

func NewJobSession(repo repository.JobSession) JobSession {
	return &jobSession{
		repo: repo,
	}
}

func (j *jobSession) List(ctx context.Context, jobID model.JobID, page model.PageNum, per model.PerPageNum, order model.Order) (model.JobSessions, error) {
	jobSessions, err := j.repo.Fetch(ctx, jobID, page, per, order)
	if err != nil {
		return nil, fmt.Errorf("failed fetch jobs page=%d, per=%d, order=%s: %w", page, per, order, err)
	}
	return jobSessions, nil
}
