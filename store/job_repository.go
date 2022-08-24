package store

import (
	"context"

	"github.com/kironono/pinkie/entity"
)

type JobRepository struct {
}

func (jr *JobRepository) List(ctx context.Context, db Queryer) (entity.Jobs, error) {
	jobs := entity.Jobs{}
	sql := `
	SELECT
		id, name, slug, created_at, updated_at
	FROM
		jobs
	;`
	if err := db.SelectContext(ctx, &jobs, sql); err != nil {
		return nil, err
	}
	return jobs, nil
}
