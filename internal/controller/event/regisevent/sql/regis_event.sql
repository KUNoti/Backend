-- name: CreateRegisEvent :one
INSERT INTO "regis_events"
(
    event_id, user_id, created_at, updated_at
)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING *;