package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
	"github.com/kironono/pinkie/store"
)

type Metric interface {
	Create(context.Context, string, time.Time, string, string) error
}

type metric struct {
	atomic         store.Atomic
	jobRepo        repository.Job
	jobSessionRepo repository.JobSession
}

func NewMetric(
	atomic store.Atomic,
	jobRepo repository.Job,
	jobSessionRepo repository.JobSession) Metric {
	return &metric{
		atomic:         atomic,
		jobRepo:        jobRepo,
		jobSessionRepo: jobSessionRepo,
	}
}

func (m *metric) Create(ctx context.Context, jobSlug string, ts time.Time, status string, task string) error {
	if err := m.atomic.DoInTx(ctx, func(ctx context.Context) error {
		job, err := m.jobRepo.GetBySlug(ctx, jobSlug)
		if err != nil {
			return fmt.Errorf("failed find job jobSlug=%s: %w", jobSlug, err)
		}

		openedSession, err := m.jobSessionRepo.GetOpenedJobSessionByJobID(ctx, job.ID)
		if err != nil {
			return fmt.Errorf("failed get opend job session jobID=%d: %w", job.ID, err)
		}

		if openedSession == nil && status == "finished" {
			return model.ErrOpendJobSessionNotFound
		}

		if openedSession != nil {
			if err := m.jobSessionRepo.CloseJobSession(ctx, openedSession.ID, ts); err != nil {
				return fmt.Errorf("failed close job session jobSessionID=%d: %w", openedSession.ID, err)
			}
		}

		if status == "started" {
			if err := m.jobSessionRepo.OpenJobSession(ctx, job.ID, ts); err != nil {
				return fmt.Errorf("failed open job session jobID=%d: %w", job.ID, err)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
