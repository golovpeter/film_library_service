-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL PRIMARY KEY,
    username      VARCHAR(20) UNIQUE NOT NULL,
    password_hash TEXT               NOT NULL,
    role          varchar(20)        NOT NULL default 'user',
    created_at    TIMESTAMP WITH TIME ZONE    DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS films
(
    id           BIGSERIAL PRIMARY KEY,
    title        VARCHAR(150) NOT NULL,
    description  VARCHAR(1000),
    release_date timestamp    not null,
    rating       DECIMAL(3, 1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE IF NOT EXISTS actors
(
    id          BIGSERIAL PRIMARY KEY,
    first_name  VARCHAR(50)              NOT NULL,
    second_name VARCHAR(50)              NOT NULL,
    birth_date  TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS files_and_actors
(
    id       BIGSERIAL,
    film_id  INT NOT NULL REFERENCES films,
    actor_id INT NOT NULL REFERENCES actors
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS files_and_actors;
-- +goose StatementEnd
