CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,

    username VARCHAR(100)
        NOT NULL
        UNIQUE,

    password_hash TEXT
        NOT NULL,

    mfa_enabled BOOLEAN
        NOT NULL
        DEFAULT FALSE,

    mfa_secret TEXT,

    failed_attempts INTEGER
        NOT NULL
        DEFAULT 0,

    locked_until TIMESTAMP,

    created_at TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    last_login TIMESTAMP
);

CREATE INDEX idx_users_username
ON users(username);