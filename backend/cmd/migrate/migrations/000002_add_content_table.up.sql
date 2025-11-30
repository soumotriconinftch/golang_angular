CREATE TABLE IF NOT EXISTS content(
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id bigint NOT NULL,
    body text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
     CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);