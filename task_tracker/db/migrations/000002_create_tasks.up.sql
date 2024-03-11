create table if not exists tasks
(
    id          uuid primary key      default uuid_generate_v4(),
    user_id     uuid references users (id),
    title       text         not null,
    description text         not null,
    status      varchar(256) not null,
    created_at  timestamp    not null default current_timestamp,
    updated_at  timestamp    not null default current_timestamp
);

create index if not exists tasks_status_idx on tasks (status);

create index if not exists tasks_created_at_idx on tasks (created_at);

create index if not exists tasks_updated_at_idx on tasks (updated_at);