package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kironono/pinkie/config"
)

func NewDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	url, err := databaseURL(cfg)
	if err != nil {
		return nil, nil, err
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() { db.Close() }, err
	}

	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { db.Close() }, nil
}

func databaseURL(cfg *config.Config) (string, error) {
	d, err := mysql.ParseDSN(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		return "", err
	}

	d.Loc = time.UTC
	d.ParseTime = true
	d.Collation = "utf8mb4_general_ci"
	if d.Params == nil {
		d.Params = map[string]string{}
	}
	d.InterpolateParams = true
	// enforce kamipo TRADITIONAL!
	d.Params["sql_mode"] = "'TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY'"
	return d.FormatDSN(), nil
}

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Preparer interface {
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Queryer interface {
	Preparer
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

var (
	_ Beginner = (*sqlx.DB)(nil)
	_ Preparer = (*sqlx.DB)(nil)
	_ Queryer  = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.DB)(nil)
	_ Execer   = (*sqlx.Tx)(nil)
)
