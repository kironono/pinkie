package repository

import (
	"context"
	"time"

	"github.com/kironono/pinkie/model"
)

type Metric interface {
	GetJobBySlug(context.Context, string) (*model.Job, error)
	GetOpenedJobSessionByJobID(context.Context, model.JobID) (*model.JobSession, error)
	CloseJobSession(context.Context, model.JobSessionID, time.Time) error
	OpenJobSession(context.Context, model.JobID, time.Time) error
}
