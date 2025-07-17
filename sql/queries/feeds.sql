-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
select feeds.name as name, url, users.name as user from feeds
join users on users.id = feeds.user_id;

-- name: GetFeed :one
select * from feeds where name = $1;

-- name: MarkFeedFetched :one
update feeds
set updated_at = current_timestamp, last_fetched_at = current_timestamp
where id = $1
returning *;

-- name: GetNextFeedToFetch :one
select * from feeds order by last_fetched_at nulls first limit 1;
