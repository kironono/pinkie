package repository

import (
	"context"
	"time"

	"github.com/kironono/pinkie/model"
)

type JobSession interface {
	Fetch(context.Context, model.JobID, model.PageNum, model.PerPageNum, model.Order) (model.JobSessions, error)
	GetOpenedJobSessionByJobID(context.Context, model.JobID) (*model.JobSession, error)
	CloseJobSession(context.Context, model.JobSessionID, time.Time) error
	OpenJobSession(context.Context, model.JobID, time.Time) error
}
