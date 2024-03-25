-- name: CreateUser :one
INSERT INTO users
(name, username, password, social_id, email, profile_image, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING *;