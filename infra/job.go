package infra

import (
	"context"

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

func (j *JobRepository) Find(id entity.JobID) (*entity.Job, error) {
	// TODO
	ctx := context.TODO()
	jobs := entity.Jobs{}
	sql := `
	SELECT
		id, name, slug, created_at, updated_at
	FROM
		jobs
	WHERE
		id = ?
	;`
	if err := j.DB.SelectContext(ctx, &jobs, sql, id); err != nil {
		return nil, err
	}
	return jobs[0], nil
}

func (j *JobRepository) Create(job *entity.Job) (*entity.Job, error) {
	// TODO:
	return job, nil
}
