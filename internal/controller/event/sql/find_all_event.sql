-- name: FindAllEvent :many
SELECT *
FROM events
WHERE regis_max > regis_amount
  AND start_date > CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok'
ORDER BY id DESC;
