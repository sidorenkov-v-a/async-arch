create table if not exists tasks
(
    id          uuid primary key      default uuid_generate_v4(),
    reporter_id uuid references users (id),
    assignee_id uuid references users (id),
    jira_id     serial,
    title       text         not null,
    description text         not null,
    status      varchar(256) not null,
    cost        integer      not null,
    created_at  timestamp    not null default current_timestamp,
    updated_at  timestamp    not null default current_timestamp
);

create index if not exists tasks_jira_idx on tasks (jira_id);

create index if not exists tasks_status_idx on tasks (status);

create index if not exists tasks_created_at_idx on tasks (created_at);

create index if not exists tasks_updated_at_idx on tasks (updated_at);

create index if not exists tasks_cost_idx on tasks (cost);