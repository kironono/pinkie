package repository

import (
	"context"

	"github.com/kironono/pinkie/model"
)

type User interface {
	GetByID(context.Context, model.UserID) (*model.User, error)
	Fetch(context.Context, model.PageNum, model.PerPageNum, model.Order) (model.Users, error)
}
