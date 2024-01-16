// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: create_test.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTest = `-- name: CreateTest :one
INSERT INTO "tests" (
    title,
    created_at,
    updated_at
)
VALUES (
    $1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING id, title, created_at, updated_at
`

func (q *Queries) CreateTest(ctx context.Context, title pgtype.Text) (Test, error) {
	row := q.db.QueryRow(ctx, createTest, title)
	var i Test
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
