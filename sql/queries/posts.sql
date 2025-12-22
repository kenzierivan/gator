-- name: CreatePost :exec
insert into posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
values($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetPosts :many
select * 
from posts
order by published_at desc
limit $1;