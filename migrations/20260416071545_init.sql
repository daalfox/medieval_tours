-- +goose Up
CREATE TABLE api_order (
    id BIGSERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS api_order;
