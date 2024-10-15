-- name: CreateSession :one
INSERT INTO sessions (
  id, user_id, email, provider, access_token, refresh_token, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSession :one
SELECT refresh_token FROM sessions
WHERE email=?
LIMIT 1;

-- name: UpdateSession :exec
UPDATE sessions
SET 
  refresh_token=?,
  access_token=?
WHERE email=?;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE email=?;
