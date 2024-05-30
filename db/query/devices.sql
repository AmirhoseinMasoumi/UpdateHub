-- name: CreateDevice :one
INSERT INTO devices (
  device_id
) VALUES (
  $1
) 
RETURNING *;

-- name: UpdateDevice :one
UPDATE devices
SET
  device_version = $1,
  last_update = $2
WHERE device_id = $3
RETURNING *;

-- name: GetDevice :one
SELECT * FROM devices WHERE device_id = $1 LIMIT 1;

-- name: ListAllDevices :many
SELECT device_id, device_version, created_at, last_update
FROM devices;

-- name: DeleteDevice :one
DELETE FROM devices 
WHERE device_id = $1
RETURNING *;