package usecase

import (
	"context"

	"github.com/kironono/pinkie/entity"
	"github.com/kironono/pinkie/repository"
)

type Job interface {
	Show(context.Context, entity.JobID) (*entity.Job, error)
	List(context.Context) (entity.Jobs, error)
}

type job struct {
	repo repository.Job
}

func NewJob(repo repository.Job) Job {
	return &job{
		repo: repo,
	}
}

func (j *job) Show(ctx context.Context, id entity.JobID) (*entity.Job, error) {
	return j.repo.First(ctx, id)
}

func (j *job) List(ctx context.Context) (entity.Jobs, error) {
	return j.repo.Find(ctx)
}
