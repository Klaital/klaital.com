
-- name: CreateUser :one
INSERT INTO users (username, email, password_digest) VALUES ($1, $2, $3) RETURNING id;

-- name: GetUserById :one
SELECT id, username, email, password_digest FROM users WHERE id = $1;

-- name: GetUser :one
SELECT id, username, email, password_digest FROM users WHERE email = $1;

-- name: UpdateUsername :exec
UPDATE users SET username = $2 WHERE id = $1;

-- name: UpdateEmail :exec
UPDATE users SET email = $2 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
