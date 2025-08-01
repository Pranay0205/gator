-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users 
WHERE name = $1 LIMIT 1;


-- name: TruncateUser :exec
TRUNCATE TABLE users;


-- name: GetUsers :many
SELECT * FROM users LIMIT 100;