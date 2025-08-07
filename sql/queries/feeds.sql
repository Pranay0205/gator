-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
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
SELECT
  f.name AS feed_name,
  f.url,
  u.name AS user_name
FROM
  feeds AS f
INNER JOIN
  users AS u ON f.user_id = u.id;


-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id FROM feeds WHERE url = $1 LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET last_fetched_at = $1, updated_at = $2 
WHERE id = $3;


-- name: GetNextFeedToFetch :one
SELECT 
  f.id, f.name, f.url, f.last_fetched_at
FROM 
  feeds f
  ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
