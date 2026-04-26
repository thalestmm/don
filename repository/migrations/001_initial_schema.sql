-- Buckets are the main aggregators for financial data
CREATE TABLE buckets (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB
);

CREATE TABLE droplets (
    id UUID PRIMARY KEY,
    bucket_id UUID NOT NULL REFERENCES buckets(id),
    name TEXT NOT NULL,
    initial_balance_cents INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB
);

CREATE TABLE drips (
    id UUID PRIMARY KEY,
    droplet_id UUID NOT NULL REFERENCES droplets(id),
    increases BOOLEAN NOT NULL DEFAULT true,
    amount_cents INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB
);
