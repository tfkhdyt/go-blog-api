-- name: CreateUser :one
INSERT INTO "user" (
  full_name, username, email, password, role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, full_name, username, email, role, created_at;

-- name: FindAllUsers :many
SELECT id, full_name, username, email, role, created_at, updated_at FROM "user";

-- name: FindOneUserByID :one
SELECT * FROM "user" 
WHERE id = $1 LIMIT 1;

-- name: FindOneUserByEmail :one
SELECT * FROM "user" 
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE "user"
SET 
  full_name = $2,
  username = $3,
  updated_at = $4
WHERE id = $1
RETURNING id, full_name, username, email, role, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
