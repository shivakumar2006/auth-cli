CREATE TABLE IF NOT EXISTS sessions (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT
        NOT NULL,

    token UUID
        NOT NULL
        UNIQUE,

    created_at TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    expires_at TIMESTAMP
        NOT NULL,

    CONSTRAINT fk_sessions_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_sessions_token
ON sessions(token);

CREATE INDEX idx_sessions_user_id
ON sessions(user_id);

CREATE INDEX idx_sessions_expires_at
ON sessions(expires_at);