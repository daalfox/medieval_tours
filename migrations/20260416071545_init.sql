-- +goose Up
CREATE TABLE tour (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);
CREATE TABLE schedule (
    id BIGSERIAL PRIMARY KEY,
    tour_id BIGINT REFERENCES tour(id),
    starts_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS tour;
