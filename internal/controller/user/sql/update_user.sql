-- name: UpdateUserByID :one
UPDATE users SET
                 name = COALESCE($1, name),
                 social_id = COALESCE($2, social_id),
                 email = COALESCE($3, email),
                 profile_image = COALESCE($4, profile_image),
                 updated_at = CURRENT_TIMESTAMP
WHERE id = $5
RETURNING *;