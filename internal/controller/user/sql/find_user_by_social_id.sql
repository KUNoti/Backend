-- name: FindUserBySocialId :one
SELECT * FROM users
WHERE social_id = $1 LIMIT 1;
