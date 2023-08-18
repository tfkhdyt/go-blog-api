-- name: AddResetPasswordToken :one
INSERT INTO reset_password_token (
  token, user_id, expires_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: FindResetPasswordToken :one
SELECT * FROM reset_password_token 
WHERE token = $1;

-- name: RemoveResetPasswordToken :exec
DELETE FROM reset_password_token
WHERE token = $1;
