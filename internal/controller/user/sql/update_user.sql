-- name: UpdateUserByID :one
UPDATE users SET
                 name = COALESCE(sqlc.narg('name'), name),
                 social_id = COALESCE(sqlc.narg('social_id'), social_id),
                 email = COALESCE(sqlc.narg('email'), email),
                 profile_image = COALESCE(sqlc.narg('profile_image'), profile_image),
                 updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;