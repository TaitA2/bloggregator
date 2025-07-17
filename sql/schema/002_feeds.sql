-- +goose Up
create table feeds (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    last_fetched_at timestamp,
    name text unique not null,
    url text unique not null,
    user_id uuid not null references users on delete cascade,
    constraint fk_user_id foreign key(user_id) references users (id)
);

-- +goose Down
drop table feeds;
