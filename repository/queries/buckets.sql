-- name: GetBuckets :many
SELECT * FROM buckets;

-- name: CreateBucket :one
INSERT INTO buckets (id, name, kind, metadata)
VALUES ($1, $2, $3, $4) RETURNING *;
