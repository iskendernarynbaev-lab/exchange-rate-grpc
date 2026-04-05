-- +goose Up
CREATE TABLE IF NOT EXISTS rates (
    id BIGSERIAL PRIMARY KEY,
    ask DOUBLE PRECISION NOT NULL,
    bid DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_rates_created_at ON rates(created_at DESC);

-- +goose Down
DROP TABLE IF EXISTS rates;
