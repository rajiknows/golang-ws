-- name: CreateUser :one
INSERT INTO users (
  id,created_at, updated_at , name, email, password
) VALUES ( $1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetPassByEmail :one
SELECT password FROM users WHERE email = $1;



