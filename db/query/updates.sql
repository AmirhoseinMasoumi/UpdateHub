-- name: CreateUpdate :one
INSERT INTO updates (
  version, path, description, checksum
) VALUES (
  $1, $2, $3, $4
) 
ON CONFLICT (version) DO UPDATE
SET path = EXCLUDED.path,
    description = EXCLUDED.description
RETURNING *;

-- name: GetUpdate :one
SELECT * FROM updates WHERE version = $1 LIMIT 1;

-- name: GetNextVersion :one
SELECT * FROM updates
WHERE version > $1
ORDER BY version ASC
LIMIT 1;

-- name: ListUpdatesBetweenDates :many
SELECT *
FROM updates
WHERE date BETWEEN sqlc.arg(start_date) AND sqlc.arg(end_date);

-- name: DeleteUpdate :one
DELETE FROM updates 
WHERE version = $1
RETURNING *;
