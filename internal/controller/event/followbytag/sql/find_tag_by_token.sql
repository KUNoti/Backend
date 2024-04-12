-- name: FindTagNameByTokens :many
SELECT tag
FROM "follow_by_tag"
WHERE user_token = $1;