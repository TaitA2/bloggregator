-- +goose Up
create table posts (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    url text unique not null,
    description text,
    published_at timestamp not null,
    feed_id uuid not null references feeds on delete cascade,
    constraint fk_feed_id foreign key(feed_id) references feeds (id)
);

-- +goose Down
drop table posts;
