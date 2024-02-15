-- +goose Up
CREATE TABLE vendor (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    business_name TEXT NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone_number VARCHAR(25) UNIQUE NOT NULL,
    country VARCHAR(70) NOT NULL,
    profile_image TEXT,
    banner_image TEXT,
    description TEXT,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE vendor;