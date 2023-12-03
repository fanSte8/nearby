CREATE TABLE IF NOT EXISTS likes_new (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    post_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, post_id)
);

INSERT INTO likes_new (user_id, post_id, created_at, updated_at)
SELECT user_id, post_id, created_at, updated_at FROM likes;

DROP TABLE likes;

ALTER TABLE likes_new RENAME TO likes;
