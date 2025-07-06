-- +goose Up
create table feed_follows (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid not null references users on delete cascade,
    feed_id uuid not null references feeds on delete cascade,
    constraint fk_user_id foreign key(user_id) references users (id),
    constraint fk_feed_id foreign key(feed_id) references feeds (id),
    constraint user_feed_pair unique (user_id, feed_id)
);

-- +goose Down
drop table feed_follows;

