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
    rating       DECIMAL(3, 1) CHECK (rating >= 0 AND rating <= 10),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS actors
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(50) NOT NULL,
    gender     gender      NOT NULL,
    birth_date TIMESTAMP   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS films_and_actors
(
    id       BIGSERIAL,
    film_id  INT NOT NULL REFERENCES films,
    actor_id INT NOT NULL REFERENCES actors
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS gender;
DROP TABLE IF EXISTS films_and_actors;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS films;
DROP TABLE IF EXISTS actors;
-- +goose StatementEnd
