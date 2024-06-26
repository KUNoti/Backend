// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_noti_by_token.sql

package sqlc

import (
	"context"
)

const findNotiByToken = `-- name: FindNotiByToken :many
SELECT id, title, body, data, token, created_at, updated_at
FROM notifications
WHERE token = $1
ORDER BY id DESC
`

func (q *Queries) FindNotiByToken(ctx context.Context, token string) ([]Notification, error) {
	rows, err := q.db.Query(ctx, findNotiByToken, token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.Data,
			&i.Token,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
