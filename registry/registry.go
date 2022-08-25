package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/infra"
	"github.com/kironono/pinkie/repository"
)

type Repository interface {
	NewJob() repository.Job
}

type repositoryImpl struct {
	DB      *sqlx.DB
	jobRepo repository.Job
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		DB: db,
	}
}

func (r *repositoryImpl) NewJob() repository.Job {
	if r.jobRepo == nil {
		r.jobRepo = infra.NewJobRepository(r.DB)
	}
	return r.jobRepo
}
