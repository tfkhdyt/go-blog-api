-- name: AddRefreshToken :one
INSERT INTO refresh_token (
  token, user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: FindRefreshToken :one
SELECT * FROM refresh_token 
WHERE token = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_token
WHERE token = $1;
