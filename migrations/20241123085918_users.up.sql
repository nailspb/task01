--CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users
(
    --id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    id         SERIAL PRIMARY KEY ,
    email      varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP          DEFAULT NULL
);