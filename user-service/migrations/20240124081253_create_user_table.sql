-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    uuid string NOT NULL,
    name string NOT NULL,
    surname string,
    email string NOT NULL,
    hash_password string NOT NULL,
    role string,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
