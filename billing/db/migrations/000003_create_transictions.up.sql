create table if not exists billing_cycles
(
    id          uuid primary key   default uuid_generate_v4(),
    started_at  timestamp not null default current_timestamp,
    finished_at timestamp null
);

create table if not exists transactions
(
    id               uuid primary key default uuid_generate_v4(),
    user_id          uuid   not null references users (id),
    billing_cycle_id uuid   not null references billing_cycles (id),
    description      text   not null,
    debit            bigint null,
    credit           bigint null,
    type             varchar(256)
)