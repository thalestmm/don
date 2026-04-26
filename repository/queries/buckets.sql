-- name: GetBuckets :many
SELECT * FROM buckets;

-- name: CreateBucket :one
INSERT INTO buckets (id, name, metadata)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCurrentBucketBalance :one
SELECT
    COALESCE((
        SELECT SUM(
            CASE
                WHEN d.increases = true THEN d.amount_cents
                ELSE -d.amount_cents
            END
        )
        FROM droplets d
        WHERE d.bucket_id = $1
    ), 0) AS total_bucket_balance
WHERE d.bucket_id = $1;
