-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE department(
    id BIGINT NOT NULL,
    name VARCHAR(256),
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp)
);

CREATE TABLE book(
    id BIGINT NOT NULL,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year INT,
    price VARCHAR(256),
    total_page INT,
    thickness VARCHAR(256),
    created_at TIMESTAMP default(current_timestamp),
    updated_at TIMESTAMP default(current_timestamp),
    category_id BIGINT NOT NULL
)

-- +migrate StatementEnd