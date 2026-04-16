-- +goose Up
CREATE TABLE tour (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS tour;
