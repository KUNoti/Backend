// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: unfollow_event.sql

package sqlc

import (
	"context"
)

const unfollowEvent = `-- name: UnfollowEvent :one
DELETE FROM "following_events" WHERE event_id = $1 AND user_id = $2
RETURNING event_id
`

type UnfollowEventParams struct {
	EventID int32 `json:"event_id"`
	UserID  int32 `json:"user_id"`
}

func (q *Queries) UnfollowEvent(ctx context.Context, arg UnfollowEventParams) (int32, error) {
	row := q.db.QueryRow(ctx, unfollowEvent, arg.EventID, arg.UserID)
	var event_id int32
	err := row.Scan(&event_id)
	return event_id, err
}
