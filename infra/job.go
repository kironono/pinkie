package infra

import (
	"context"
	"database/sql"
	"fmt"

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

func (j *JobRepository) GetByID(ctx context.Context, id model.JobID) (*model.Job, error) {
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

func (j *JobRepository) Fetch(ctx context.Context, page model.PageNum, per model.PerPageNum, order model.Order) (model.Jobs, error) {
	jobs := model.Jobs{}
	q := fmt.Sprintf(`SELECT * FROM jobs ORDER BY %s LIMIT ? OFFSET ?`, order)

	offset := (int(page) - 1) * int(per)
	limit := int(per)

	if err := j.DB.SelectContext(ctx, &jobs, q, limit, offset); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobRepository) Store(ctx context.Context, job *model.Job) error {
	q := `INSERT jobs SET name=?, slug=?, created_at=?, updated_at=?`
	stmt, err := j.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, job.Name, job.Slug, job.CreatedAt, job.UpdatedAt)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	job.ID = model.JobID(lastID)
	return nil
}

func (j *JobRepository) Update(ctx context.Context, job *model.Job) error {
	q := `UPDATE jobs SET name=?, slug=?, updated_at=? WHERE id = ?`
	stmt, err := j.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, job.Name, job.Slug, job.UpdatedAt, job.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("illegal rows affected count=%d: %w", rowsAffected, err)
	}
	return nil
}

func (j *JobRepository) Delete(ctx context.Context, id model.JobID) error {
	q := `DELETE FROM jobs WHERE id = ?`
	stmt, err := j.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("illegal rows affected count=%d: %w", rowsAffected, err)
	}
	return nil
}
