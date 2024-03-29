-- name: UnfollowEvent :one
DELETE FROM "following_events" WHERE event_id = $1 AND user_id = $2
RETURNING event_id;
