-- name: FindAllEvent :many
SELECT *
FROM events
ORDER BY id DESC;