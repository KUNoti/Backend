-- name: FindAllFollowEvent :many
SELECT *
FROM following_events
WHERE user_id = $1
ORDER BY id DESC;