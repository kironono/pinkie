package repository

import "github.com/kironono/pinkie/entity"

type Job interface {
	Find(entity.JobID) (*entity.Job, error)
	Create(*entity.Job) (*entity.Job, error)
}
