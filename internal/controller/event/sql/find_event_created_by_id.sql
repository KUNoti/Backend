-- name: FindEventCreatedByID :many
SELECT *
FROM events
WHERE creator = $1
ORDER BY id DESC;