package repository

import (
	"context"

	"github.com/kironono/pinkie/model"
)

type Job interface {
	GetByID(context.Context, model.JobID) (*model.Job, error)
	Fetch(context.Context, model.PageNum, model.PerPageNum, model.Order) (model.Jobs, error)
	Store(context.Context, *model.Job) error
	Update(context.Context, *model.Job) error
	Delete(context.Context, model.JobID) error
}
