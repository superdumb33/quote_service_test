CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS quotes (
    id          UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    author      TEXT            NOT NULL,
    quote       TEXT            NOT NULL,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_quotes_author
    ON quotes(author);