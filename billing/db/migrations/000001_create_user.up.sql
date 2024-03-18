create extension if not exists "uuid-ossp";

create table if not exists users
(
    id         uuid primary key      default uuid_generate_v4(),
    email      text         not null,
    role       varchar(256) not null default 'worker',
    first_name text         not null,
    last_name  text         not null,
    balance    bigint       not null default 0,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp
);

create index if not exists users_email_idx
    ON users (email);