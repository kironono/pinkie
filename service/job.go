package service

import (
	"context"
	"fmt"

	"github.com/kironono/pinkie/entity"
	"github.com/kironono/pinkie/store"
)

type JobLister interface {
	List(ctx context.Context, db store.Queryer) (entity.Jobs, error)
}

type JobService struct {
	DB   store.Queryer
	Repo JobLister
}

func (s *JobService) ListJobs(ctx context.Context) (entity.Jobs, error) {
	jobs, err := s.Repo.List(ctx, s.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return jobs, err
}
