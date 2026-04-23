-- name: GetDrips :many
SELECT * FROM drips;

-- name: CreateDrip :one
INSERT INTO drips (droplet_id, increases, amount_cents, metadata) VALUES ($1, $2, $3, $4) RETURNING *;
