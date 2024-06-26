// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Device struct {
	DeviceID      string    `json:"device_id"`
	DeviceVersion string    `json:"device_version"`
	LastUpdate    time.Time `json:"last_update"`
	CreatedAt     time.Time `json:"created_at"`
}

type Update struct {
	Version     string    `json:"version"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	Checksum    string    `json:"checksum"`
	Date        time.Time `json:"date"`
}
