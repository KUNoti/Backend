-- name: DeleteEventByID :execresult
DELETE FROM "events" WHERE id = $1
RETURNING id;