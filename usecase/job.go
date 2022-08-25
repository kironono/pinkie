package usecase

import (
	"github.com/kironono/pinkie/entity"
	"github.com/kironono/pinkie/repository"
)

type Job interface {
	Show(entity.JobID) (*entity.Job, error)
}

type job struct {
	repo repository.Job
}

func NewJob(repo repository.Job) Job {
	return &job{
		repo: repo,
	}
}

func (j *job) Show(id entity.JobID) (*entity.Job, error) {
	return j.repo.Find(id)
}
