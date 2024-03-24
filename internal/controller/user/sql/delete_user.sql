-- name: DeleteUser :execresult
DELETE FROM users WHERE id = $1;