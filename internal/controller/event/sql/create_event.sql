-- name: CreateEvent :one
INSERT INTO "events"
(
 title, latitude, longitude, start_date, end_date, price, rating,
 image, creator, detail, location_name, need_regis,
 created_at, updated_at
 )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;