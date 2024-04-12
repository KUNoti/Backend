-- name: CreateNotification :one
INSERT INTO "notifications"
(
    body, data, title, token, created_at, updated_at
)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING *;