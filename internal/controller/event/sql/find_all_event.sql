-- name: FindAllEvent :many
SELECT *
FROM events
WHERE regis_max > regis_amount
ORDER BY id DESC;