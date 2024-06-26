// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: devices.sql

package db

import (
	"context"
	"time"
)

const createDevice = `-- name: CreateDevice :one
INSERT INTO devices (
  device_id
) VALUES (
  $1
) 
RETURNING device_id, device_version, last_update, created_at
`

func (q *Queries) CreateDevice(ctx context.Context, deviceID string) (Device, error) {
	row := q.db.QueryRowContext(ctx, createDevice, deviceID)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.DeviceVersion,
		&i.LastUpdate,
		&i.CreatedAt,
	)
	return i, err
}

const deleteDevice = `-- name: DeleteDevice :one
DELETE FROM devices 
WHERE device_id = $1
RETURNING device_id, device_version, last_update, created_at
`

func (q *Queries) DeleteDevice(ctx context.Context, deviceID string) (Device, error) {
	row := q.db.QueryRowContext(ctx, deleteDevice, deviceID)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.DeviceVersion,
		&i.LastUpdate,
		&i.CreatedAt,
	)
	return i, err
}

const getDevice = `-- name: GetDevice :one
SELECT device_id, device_version, last_update, created_at FROM devices WHERE device_id = $1 LIMIT 1
`

func (q *Queries) GetDevice(ctx context.Context, deviceID string) (Device, error) {
	row := q.db.QueryRowContext(ctx, getDevice, deviceID)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.DeviceVersion,
		&i.LastUpdate,
		&i.CreatedAt,
	)
	return i, err
}

const listAllDevices = `-- name: ListAllDevices :many
SELECT device_id, device_version, created_at, last_update
FROM devices
`

type ListAllDevicesRow struct {
	DeviceID      string    `json:"device_id"`
	DeviceVersion string    `json:"device_version"`
	CreatedAt     time.Time `json:"created_at"`
	LastUpdate    time.Time `json:"last_update"`
}

func (q *Queries) ListAllDevices(ctx context.Context) ([]ListAllDevicesRow, error) {
	rows, err := q.db.QueryContext(ctx, listAllDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAllDevicesRow{}
	for rows.Next() {
		var i ListAllDevicesRow
		if err := rows.Scan(
			&i.DeviceID,
			&i.DeviceVersion,
			&i.CreatedAt,
			&i.LastUpdate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDevice = `-- name: UpdateDevice :one
UPDATE devices
SET
  device_version = $1,
  last_update = $2
WHERE device_id = $3
RETURNING device_id, device_version, last_update, created_at
`

type UpdateDeviceParams struct {
	DeviceVersion string    `json:"device_version"`
	LastUpdate    time.Time `json:"last_update"`
	DeviceID      string    `json:"device_id"`
}

func (q *Queries) UpdateDevice(ctx context.Context, arg UpdateDeviceParams) (Device, error) {
	row := q.db.QueryRowContext(ctx, updateDevice, arg.DeviceVersion, arg.LastUpdate, arg.DeviceID)
	var i Device
	err := row.Scan(
		&i.DeviceID,
		&i.DeviceVersion,
		&i.LastUpdate,
		&i.CreatedAt,
	)
	return i, err
}
