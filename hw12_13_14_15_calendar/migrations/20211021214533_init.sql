-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table event
(
    id            uuid default uuid_generate_v4() not null
        constraint event_pkey
            primary key,
    title         varchar(256),
    "Description" text,
    "UserID"      uuid,
    "StartTime"   timestamp,
    "EndTime"     timestamp
);

alter table event owner to postgres;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
