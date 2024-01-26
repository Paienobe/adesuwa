-- +goose Up
CREATE TABLE buyer (
    id UUID PRIMARY KEY NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    profile_image TEXT,
    password TEXT NOT NULL
);

-- +goose Down
DROP TABLE buyer;