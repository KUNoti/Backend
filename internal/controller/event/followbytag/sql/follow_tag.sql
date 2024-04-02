-- name: FollowTag :one
INSERT INTO "follow_by_tag"
(
    tag, user_token, created_at, updated_at
)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING *;