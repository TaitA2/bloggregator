-- name: CreateFeedFollow :one
with inserted_feed_follow as (
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values ( $1, $2, $3, $4, $5)
returning *
)
select
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
JOIN feeds on feeds.id = inserted_feed_follow.feed_id
JOIN users on users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
select *, users.name as user, feeds.name as feed from feed_follows 
join users on users.id = user_id
join feeds on feeds.id = feed_id
where feed_follows.user_id = $1;

-- name: Unfollow :exec
delete from feed_follows
where feed_id = (select id from feeds where url = $1);
