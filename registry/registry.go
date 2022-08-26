package registry

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/infra"
	"github.com/kironono/pinkie/repository"
	"github.com/kironono/pinkie/store"
)

type Repository interface {
	NewAtomic() store.Atomic
	NewJob() repository.Job
	NewUser() repository.User
	NewMetric() repository.Metric
}

type repositoryImpl struct {
	DB         *sqlx.DB
	Atomic     store.Atomic
	jobRepo    repository.Job
	userRepo   repository.User
	metricRepo repository.Metric
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{
		DB:     db,
		Atomic: store.NewAtomic(db),
	}
}

func (r *repositoryImpl) DoInTx(ctx context.Context) error {
	r.DB.BeginTxx(ctx, nil)
	return nil
}

func (r *repositoryImpl) NewAtomic() store.Atomic {
	return r.Atomic
}

func (r *repositoryImpl) NewJob() repository.Job {
	if r.jobRepo == nil {
		r.jobRepo = infra.NewJobRepository(r.DB)
	}
	return r.jobRepo
}

func (r *repositoryImpl) NewUser() repository.User {
	if r.userRepo == nil {
		r.userRepo = infra.NewUserRepository(r.DB)
	}
	return r.userRepo
}

func (r *repositoryImpl) NewMetric() repository.Metric {
	if r.metricRepo == nil {
		r.metricRepo = infra.NewMetricRepository(r.DB)
	}
	return r.metricRepo
}
