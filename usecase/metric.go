package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type Metric interface {
	Create(context.Context, string, time.Time, string, string) error
}

type metric struct {
	repo repository.Metric
}

func NewMetric(repo repository.Metric) Metric {
	return &metric{
		repo: repo,
	}
}

func (m *metric) Create(ctx context.Context, jobSlug string, ts time.Time, status string, task string) error {
	job, err := m.repo.GetJobBySlug(ctx, jobSlug)
	if err != nil {
		return fmt.Errorf("failed find job jobSlug=%s: %w", jobSlug, err)
	}

	openedSession, err := m.repo.GetOpenedJobSessionByJobID(ctx, job.ID)
	if err != nil {
		return fmt.Errorf("failed get opend job session jobID=%d: %w", job.ID, err)
	}

	if openedSession == nil && status == "finished" {
		return model.ErrOpendJobSessionNotFound
	}

	if openedSession != nil {
		if err := m.repo.CloseJobSession(ctx, openedSession.ID, ts); err != nil {
			return fmt.Errorf("failed close job session jobSessionID=%d: %w", openedSession.ID, err)
		}
	}

	if status == "started" {
		if err := m.repo.OpenJobSession(ctx, job.ID, ts); err != nil {
			return fmt.Errorf("failed open job session jobID=%d: %w", job.ID, err)
		}
	}
	return nil
}
