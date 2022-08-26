package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/model"
	"github.com/kironono/pinkie/repository"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.User {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetByID(ctx context.Context, id model.UserID) (*model.User, error) {
	user := &model.User{}
	q := `SELECT * FROM users WHERE id = ? LIMIT 1`

	if err := u.DB.GetContext(ctx, user, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (u *UserRepository) Fetch(ctx context.Context, page model.PageNum, per model.PerPageNum, order model.Order) (model.Users, error) {
	users := model.Users{}
	q := fmt.Sprintf(`SELECT * FROM users ORDER BY %s LIMIT ? OFFSET ?`, order)

	offset := (int(page) - 1) * int(per)
	limit := int(per)

	if err := u.DB.SelectContext(ctx, &users, q, limit, offset); err != nil {
		return nil, err
	}
	return users, nil
}
