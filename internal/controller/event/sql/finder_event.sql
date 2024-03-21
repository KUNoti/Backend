-- name: FinderEvent :many
SELECT *
FROM events
WHERE
    (title LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))
    OR (creator LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))
    OR (location_name LIKE CONCAT('%', sqlc.arg('keyword')::text, '%'))

ORDER BY id DESC;