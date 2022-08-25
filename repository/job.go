package repository

import (
	"context"

	"github.com/kironono/pinkie/model"
)

type Job interface {
	First(context.Context, model.JobID) (*model.Job, error)
	Find(context.Context) (model.Jobs, error)
	Create(context.Context, *model.Job) (*model.Job, error)
}
