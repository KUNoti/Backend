-- name: FinderEvent :many
SELECT *
FROM events
WHERE
    (id = sqlc.narg('id') OR sqlc.narg('id') IS NULL)
  AND (title = sqlc.narg('title') OR sqlc.narg('title') IS NULL)
  AND (price = sqlc.narg('price') OR sqlc.narg('price') IS NULL)
  AND (creator = sqlc.narg('creator') OR sqlc.narg('creator') IS NULL)
  AND (location_name = sqlc.narg('location_name') OR sqlc.narg('location_name') IS NULL)
  AND (need_regis = sqlc.narg('need_regis') OR sqlc.narg('need_regis') IS NULL)
  AND ((start_date >= sqlc.narg('start_date') AND end_date <= sqlc.narg('end_date'))
    OR (sqlc.narg('start_date') IS NULL AND sqlc.narg('end_date') IS NULL))
ORDER BY id DESC;