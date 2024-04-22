-- name: FindEventByTime :many
SELECT *
FROM events
WHERE start_date BETWEEN
        (CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok') - interval '5 minutes'
  AND
    (CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok') + interval '5 minutes'
  AND
    need_noti = true
ORDER BY id DESC;
