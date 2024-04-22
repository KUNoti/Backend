-- name: FindUserThatRegis :many
SELECT users.id, users.token
FROM regis_events
         JOIN users ON regis_events.user_id = users.id
WHERE regis_events.event_id = $1;
