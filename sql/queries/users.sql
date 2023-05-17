-- name: Createuser :one
INSERT INTO users (id, created_at, updated_at,name) values ($1,$2,$3,$4) returning *;