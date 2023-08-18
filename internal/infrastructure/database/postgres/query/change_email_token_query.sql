-- name: AddChangeEmailToken :one
INSERT INTO change_email_token (
  token, new_email, user_id, expires_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: FindChangeEmailToken :one
SELECT * FROM change_email_token 
WHERE token = $1;

-- name: DeleteChangeEmailToken :exec
DELETE FROM change_email_token
WHERE token = $1;
