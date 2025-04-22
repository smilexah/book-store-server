CREATE TABLE books
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    isbn        VARCHAR(20)  NOT NULL,
    description TEXT,
    author      VARCHAR(255),
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP
);