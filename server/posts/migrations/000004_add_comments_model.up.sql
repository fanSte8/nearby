CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    post_id BIGINT NOT NULL REFERENCES posts ON DELETE CASCADE,
    text TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);