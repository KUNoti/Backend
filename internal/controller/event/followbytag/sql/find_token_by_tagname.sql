-- name: FindTokensByTagName :many
SELECT user_token
FROM "follow_by_tag"
WHERE tag = $1;