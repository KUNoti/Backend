-- name: UnfollowTag :exec
DELETE FROM "follow_by_tag"
WHERE tag = $1 AND user_token = $2
RETURNING tag;