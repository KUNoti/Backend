-- name: CreateTest :one
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
RETURNING *;