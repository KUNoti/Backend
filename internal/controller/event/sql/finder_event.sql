-- name: FinderEvent :many
SELECT *
FROM events
WHERE
    (title LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))
    OR (creator LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))
    OR (location_name LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))
    AND start_date > CURRENT_TIMESTAMP AT TIME ZONE 'Asia/Bangkok'

ORDER BY id DESC;