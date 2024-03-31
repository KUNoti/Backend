-- name: FindEventByID :one
SELECT *
FROM events
WHERE id = $1
ORDER BY id DESC;