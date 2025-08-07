-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;



-- name: GetPosts :many
SELECT 
  p.id,
  p.created_at,
  p.updated_at,
  p.title,
  p.url,
  p.description,
  p.published_at,
  f.name AS feed_name
FROM
  posts AS p
INNER JOIN
  feeds AS f ON p.feed_id = f.id
ORDER BY
  p.published_at DESC
LIMIT $1;