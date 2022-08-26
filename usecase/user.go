package usecase

import (
	"context"
	"fmt"

	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type User interface {
	Show(context.Context, model.UserID) (*model.User, error)
	List(context.Context, model.PageNum, model.PerPageNum, model.Order) (model.Users, error)
}

type user struct {
	repo repository.User
}

func NewUser(repo repository.User) User {
	return &user{
		repo: repo,
	}
}

func (u *user) Show(ctx context.Context, id model.UserID) (*model.User, error) {
	job, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed got user %d: %w", id, err)
	}
	return job, err
}

func (u *user) List(ctx context.Context, page model.PageNum, per model.PerPageNum, order model.Order) (model.Users, error) {
	users, err := u.repo.Fetch(ctx, page, per, order)
	if err != nil {
		return nil, fmt.Errorf("failed fetch users page=%d, per=%d, order=%s: %w", page, per, order, err)
	}
	return users, nil
}
