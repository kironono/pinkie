package usecase

import (
	"context"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type Job interface {
	Show(context.Context, model.JobID) (*model.Job, error)
	List(context.Context) (model.Jobs, error)
}

type job struct {
	repo repository.Job
}

func NewJob(repo repository.Job) Job {
	return &job{
		repo: repo,
	}
}

func (j *job) Show(ctx context.Context, id model.JobID) (*model.Job, error) {
	job, err := j.repo.First(ctx, id)
	return job, err
}

func (j *job) List(ctx context.Context) (model.Jobs, error) {
	return j.repo.Find(ctx)
}
