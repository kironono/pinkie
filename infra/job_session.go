package infra

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type JobSessionRepository struct {
	DB *sqlx.DB
}

func NewJobSessionRepository(db *sqlx.DB) repository.JobSession {
	return &JobSessionRepository{
		DB: db,
	}
}

func (j *JobSessionRepository) GetOpenedJobSessionByJobID(ctx context.Context, jobID model.JobID) (*model.JobSession, error) {
	q := `SELECT * FROM job_sessions WHERE end_at IS NULL AND job_id = ? LIMIT 1`

	jobSession := &model.JobSession{}
	if err := j.DB.GetContext(ctx, jobSession, q, jobID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return jobSession, nil
}

func (j *JobSessionRepository) CloseJobSession(ctx context.Context, jobSessionID model.JobSessionID, ts time.Time) error {
	q := `UPDATE job_sessions SET end_at=?, updated_at=? WHERE id = ?`
	stmt, err := j.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, ts, time.Now(), jobSessionID)
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

func (j *JobSessionRepository) OpenJobSession(ctx context.Context, jobID model.JobID, ts time.Time) error {
	q := `INSERT INTO job_sessions SET job_id=?, start_at=?, created_at=?, updated_at=?`
	stmt, err := j.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = stmt.ExecContext(ctx, jobID, ts, now, now)
	if err != nil {
		return err
	}
	return nil
}
