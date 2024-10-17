-- name: CreateSession :one
INSERT INTO sessions (
  id, user_id, email, provider, access_token, refresh_token, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
ON CONFLICT (email)
DO UPDATE SET 
  refresh_token = EXCLUDED.refresh_token, 
  access_token = EXCLUDED.access_token, 
  updated_at = EXCLUDED.updated_at
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE email=?
LIMIT 1;

-- name: UpdateSession :exec
UPDATE sessions
SET 
  refresh_token=?,
  access_token=?
WHERE id=?;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE email=?;
