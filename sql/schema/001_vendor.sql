-- +goose Up
CREATE TABLE vendor (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    profile_image TEXT,
    banner_image TEXT,
    description TEXT,
    password TEXT NOT NULL
);
-- +goose Down
DROP TABLE vendor;