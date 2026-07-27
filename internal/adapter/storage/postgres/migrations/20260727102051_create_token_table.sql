-- +goose Up
CREATE TABLE tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    userID UUID NOT NULL REFERENCES users(id),
    token UUID 
)

-- +goose Down
SELECT 'down SQL query';
