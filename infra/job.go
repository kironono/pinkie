package infra

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/model"
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

func (j *JobRepository) First(ctx context.Context, id model.JobID) (*model.Job, error) {
	job := &model.Job{}
	q := `SELECT * FROM jobs WHERE id = ? LIMIT 1`

	if err := j.DB.GetContext(ctx, job, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return job, nil
}

func (j *JobRepository) Find(ctx context.Context) (model.Jobs, error) {
	jobs := model.Jobs{}
	q := `SELECT * FROM jobs`

	if err := j.DB.SelectContext(ctx, &jobs, q); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobRepository) Create(ctx context.Context, job *model.Job) (*model.Job, error) {
	// TODO:
	return job, nil
}
