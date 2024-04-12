-- name: FindNotiByToken :many
SELECT *
FROM notifications
WHERE token = $1
ORDER BY id DESC;