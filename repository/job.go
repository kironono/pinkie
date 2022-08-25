package repository

import (
	"context"

	"github.com/kironono/pinkie/entity"
)

type Job interface {
	First(context.Context, entity.JobID) (*entity.Job, error)
	Find(context.Context) (entity.Jobs, error)
	Create(context.Context, *entity.Job) (*entity.Job, error)
}
