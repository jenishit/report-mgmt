-- +goose Up

-- Access tokens are stateless JWTs with no built-in revocation. This table is
-- a denylist: logout (or any future forced sign-out) inserts the session_id
-- here, and authMiddleware rejects any token whose session_id is present,
-- even if the JWT itself hasn't expired yet.
CREATE TABLE revoked_sessions (
    session_id UUID PRIMARY KEY,
    revoked_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS revoked_sessions;
