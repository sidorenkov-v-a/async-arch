create table if not exists transactions
(
    id               uuid primary key default uuid_generate_v4(),
    user_id          uuid   not null references users (id),
    description      text   not null,
    debit            bigint null,
    credit           bigint null
)

