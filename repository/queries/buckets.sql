-- name: GetBuckets :many
SELECT * FROM buckets;

-- name: CreateBucket :one
INSERT INTO buckets (id, name, metadata)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCurrentBucketBalance :one
SELECT
    COALESCE(SUM(d.initial_balance_cents), 0) +
    COALESCE((
        SELECT SUM(
            CASE
                WHEN dr.increases = true THEN dr.amount_cents
                ELSE -dr.amount_cents
            END
        )
        FROM drips dr
        JOIN droplets inner_d ON dr.droplet_id = inner_d.id
        WHERE inner_d.bucket_id = $1
    ), 0) AS total_bucket_balance
FROM droplets d
WHERE d.bucket_id = $1;
