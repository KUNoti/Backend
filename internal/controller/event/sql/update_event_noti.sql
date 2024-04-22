-- name: UpdateEventNoti :one
UPDATE events SET
                  need_noti = false,
                  updated_at = CURRENT_TIMESTAMP


WHERE id = sqlc.arg('id')
    RETURNING *;