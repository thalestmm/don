-- name: GetDroplets :many
SELECT * FROM droplets;

-- name: GetDropletsByBucket :many
SELECT * FROM droplets WHERE bucket_id = $1;

-- name: GetDropletById :one
SELECT * FROM droplets WHERE id = $1;

-- name: CreateDroplet :one
INSERT INTO droplets (bucket_id, name, initial_balance_cents) VALUES ($1, $2, $3) RETURNING *;

-- name: GetCurrentDropletBalance :one
WITH drip_totals AS (
    SELECT
        COALESCE(SUM(amount_cents) FILTER (WHERE increases = true), 0) AS pos_var,
        COALESCE(SUM(amount_cents) FILTER (WHERE increases = false), 0) AS neg_var
    FROM drips
    WHERE droplet_id = $1
)
SELECT d.initial_balance_cents + t.pos_var - t.neg_var AS current_balance
FROM droplets d
CROSS JOIN drip_totals t
WHERE d.id = $1;
