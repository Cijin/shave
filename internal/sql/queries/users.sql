-- name: CreateUser :one
INSERT INTO users (
  id, email, sub, name, email_verified, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email=? 
LIMIT 1;
