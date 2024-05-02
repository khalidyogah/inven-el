-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE users (
    username VARCHAR(256) PRIMARY KEY,
    password VARCHAR(256)
);

CREATE TABLE department(
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp)
);

CREATE TABLE item_condition(
    id SERIAL PRIMARY KEY,
    item_id BIGINT NOT NULL,
    image_url VARCHAR(256),
    lastcheck_date TIMESTAMP,
    overall_conditon INT,
    status VARCHAR(256),
    note VARCHAR(256),
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp)
);

CREATE TABLE item_list(
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    bought_date TIMESTAMP,
    image_url VARCHAR(256),
    department_id BIGINT NOT NULL,
    type_id BIGINT NOT NULL,
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp)
);

CREATE TABLE item_type(
    id SERIAL PRIMARY KEY,
    type VARCHAR(256),
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp)
);
-- +migrate StatementEnd