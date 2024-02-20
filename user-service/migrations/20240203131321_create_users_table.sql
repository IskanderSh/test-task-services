-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uuid VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    hash_password VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    role VARCHAR(15) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
