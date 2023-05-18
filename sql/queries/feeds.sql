-- name: Createfeeds :one
INSERT INTO feeds (id, created_at, updated_at, name,url, user_id) values ($1,$2,$3,$4,$5,$6) returning *;

-- name: Getfeeds :many
SELECT * FROM feeds ;

