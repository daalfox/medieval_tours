-- +goose Up
CREATE TABLE tour (
    id BIGSERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS tour;
