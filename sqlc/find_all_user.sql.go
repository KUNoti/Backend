// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: find_all_user.sql

package sqlc

import (
	"context"
)

const findAllUser = `-- name: FindAllUser :many
SELECT id, name, created_at, updated_at, email, profile_image, username, password, social_id, token FROM users
`

func (q *Queries) FindAllUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, findAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Email,
			&i.ProfileImage,
			&i.Username,
			&i.Password,
			&i.SocialID,
			&i.Token,
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
