-- +goose Up
-- +goose StatementBegin
CREATE TABLE items
(
    id    VARCHAR PRIMARY KEY,
    name  VARCHAR UNIQUE NOT NULL,
    price INTEGER        NOT NULL
);

CREATE TABLE pickup_points
(
    id      VARCHAR PRIMARY KEY,
    address VARCHAR UNIQUE NOT NULL
);

CREATE TABLE payments
(
    id   VARCHAR PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS pickup_points;
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
