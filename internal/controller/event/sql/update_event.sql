-- name: UpdateEventByID :one
UPDATE events SET
title = COALESCE(sqlc.narg('title'), title),
latitude = COALESCE(sqlc.narg('latitude'), latitude),
longitude = COALESCE(sqlc.narg('longitude'), longitude),
start_date = COALESCE(sqlc.narg('start_date'), start_date),
end_date = COALESCE(sqlc.narg('end_date'), end_date),
price = COALESCE(sqlc.narg('price'), price),
tag = COALESCE(sqlc.narg('tag'), tag),
regis_amount = COALESCE(sqlc.narg('regis_amount'), regis_amount),
regis_max = COALESCE(sqlc.narg('regis_max'), regis_max),
creator = COALESCE(sqlc.narg('creator'), creator),
detail = COALESCE(sqlc.narg('detail'), detail),
location_name = COALESCE(sqlc.narg('location_name'), location_name),
need_regis = COALESCE(sqlc.narg('need_regis'), need_regis),
updated_at = CURRENT_TIMESTAMP


WHERE id = sqlc.arg('id')
RETURNING *;