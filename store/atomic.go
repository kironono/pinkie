package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type InTxFunc func(context.Context) error

type Atomic interface {
	DoInTx(context.Context, InTxFunc) error
}

type atomic struct {
	DB *sqlx.DB
}

func NewAtomic(db *sqlx.DB) Atomic {
	return &atomic{
		DB: db,
	}
}

func (a *atomic) DoInTx(ctx context.Context, f InTxFunc) error {
	tx, err := a.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	if err := f(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return nil
}
