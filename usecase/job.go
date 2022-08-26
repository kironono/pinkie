package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type Job interface {
	Show(context.Context, model.JobID) (*model.Job, error)
	List(context.Context, model.PageNum, model.PerPageNum, model.Order) (model.Jobs, error)
	Create(context.Context, string, string) (*model.Job, error)
	Update(context.Context, model.JobID, string, string) (*model.Job, error)
	Delete(context.Context, model.JobID) error
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
	job, err := j.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed got job %d: %w", id, err)
	}
	return job, err
}

func (j *job) List(ctx context.Context, page model.PageNum, per model.PerPageNum, order model.Order) (model.Jobs, error) {
	jobs, err := j.repo.Fetch(ctx, page, per, order)
	if err != nil {
		return nil, fmt.Errorf("failed fetch jobs page=%d, per=%d, order=%s: %w", page, per, order, err)
	}
	return jobs, nil
}

func (j *job) Create(ctx context.Context, name string, slug string) (*model.Job, error) {
	now := time.Now()
	job := &model.Job{
		Name:      name,
		Slug:      slug,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := j.repo.Store(ctx, job); err != nil {
		return nil, fmt.Errorf("faield store job %v: %w", job, err)
	}
	return job, nil
}

func (j *job) Update(ctx context.Context, id model.JobID, name string, slug string) (*model.Job, error) {
	job, err := j.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed got job %d: %w", id, err)
	}
	job.Name = name
	job.Slug = slug
	job.UpdatedAt = time.Now()
	if err := j.repo.Update(ctx, job); err != nil {
		return nil, fmt.Errorf("failed update job %v: %w", job, err)
	}
	return job, nil
}

func (j *job) Delete(ctx context.Context, id model.JobID) error {
	job, err := j.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed got job %d: %w", id, err)
	}
	if err := j.repo.Delete(ctx, job.ID); err != nil {
		return fmt.Errorf("failed delete job %v: %w", job, err)
	}
	return nil
}
