-- name: FindRegisEventByUserID :many
SELECT *
FROM regis_events
WHERE user_id = $1
ORDER BY id DESC;