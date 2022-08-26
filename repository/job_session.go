package repository

import (
	"context"
	"time"

	"github.com/kironono/pinkie/model"
)

type JobSession interface {
	GetOpenedJobSessionByJobID(context.Context, model.JobID) (*model.JobSession, error)
	CloseJobSession(context.Context, model.JobSessionID, time.Time) error
	OpenJobSession(context.Context, model.JobID, time.Time) error
}
