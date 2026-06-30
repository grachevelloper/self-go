-- +goose Up
CREATE TYPE book_status AS ENUM ('reading', 'in_wishlist', 'finished');

CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    status book_status NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS books;

DROP TYPE book_status;
