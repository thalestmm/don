-- name: GetDroplets :many
SELECT * FROM droplets;

-- name: GetDropletsByBucket :many
SELECT * FROM droplets WHERE bucket_id = $1;

-- name: GetDropletById :one
SELECT * FROM droplets WHERE id = $1;

-- name: CreateDroplet :one
INSERT INTO droplets (bucket_id, name, initial_balance_cents) VALUES ($1, $2, $3) RETURNING *;
