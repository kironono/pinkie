package infra

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/entity"
	"github.com/kironono/pinkie/repository"
)

type JobRepository struct {
	DB *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) repository.Job {
	return &JobRepository{
		DB: db,
	}
}

func (j *JobRepository) First(ctx context.Context, id entity.JobID) (*entity.Job, error) {
	job := &entity.Job{}
	q := `SELECT * FROM jobs WHERE id = ? LIMIT 1`

	if err := j.DB.GetContext(ctx, job, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		} else {
			return nil, err
		}
	}
	return job, nil
}

func (j *JobRepository) Find(ctx context.Context) (entity.Jobs, error) {
	jobs := entity.Jobs{}
	q := `SELECT * FROM jobs`

	if err := j.DB.SelectContext(ctx, &jobs, q); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobRepository) Create(ctx context.Context, job *entity.Job) (*entity.Job, error) {
	// TODO:
	return job, nil
}
