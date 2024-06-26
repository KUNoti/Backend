// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: unfollow_tag.sql

package sqlc

import (
	"context"
)

const unfollowTag = `-- name: UnfollowTag :exec
DELETE FROM "follow_by_tag"
WHERE tag = $1 AND user_token = $2
RETURNING tag
`

type UnfollowTagParams struct {
	Tag       string `json:"tag"`
	UserToken string `json:"user_token"`
}

func (q *Queries) UnfollowTag(ctx context.Context, arg UnfollowTagParams) error {
	_, err := q.db.Exec(ctx, unfollowTag, arg.Tag, arg.UserToken)
	return err
}
