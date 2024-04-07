-- name: RegisEventByID :one
UPDATE events
SET regis_amount = CASE
                       WHEN regis_max > regis_amount THEN regis_amount + 1
                       ELSE regis_amount
    END,
    updated_at = CASE
                     WHEN regis_max > regis_amount THEN CURRENT_TIMESTAMP
                     ELSE updated_at
    END
WHERE id = sqlc.arg('id')
      RETURNING *;
