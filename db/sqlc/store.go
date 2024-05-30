// store.go
package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// func NewStore(querier Querier) *Store {
// 	return &Store{querier: querier}
// }

// func (s *Store) CreateDevice(ctx context.Context, deviceID string) (Device, error) {
// 	return s.querier.CreateDevice(ctx, deviceID)
// }

// func (s *Store) CreateUpdate(ctx context.Context, arg CreateUpdateParams) (Update, error) {
// 	return s.querier.CreateUpdate(ctx, arg)
// }

// func (s *Store) DeleteDevice(ctx context.Context, deviceID string) (Device, error) {
// 	return s.querier.DeleteDevice(ctx, deviceID)
// }

// func (s *Store) DeleteUpdate(ctx context.Context, version string) (Update, error) {
// 	return s.querier.DeleteUpdate(ctx, version)
// }

// func (s *Store) UpdateDevice(ctx context.Context, arg UpdateDeviceParams) (Device, error) {
// 	return s.querier.UpdateDevice(ctx, arg)
// }

// func (s *Store) ListAllDevices(ctx context.Context) ([]ListAllDevicesRow, error) {
// 	return s.querier.ListAllDevices(ctx)
// }

// func (s *Store) ListUpdatesBetweenDates(ctx context.Context, startDate, endDate time.Time) ([]Update, error) {
// 	params := ListUpdatesBetweenDatesParams{
// 		StartDate: startDate,
// 		EndDate:   endDate,
// 	}
// 	return s.querier.ListUpdatesBetweenDates(ctx, params)
// }

// func (s *Store) GetNextVersion(ctx context.Context, version string) (Update, error) {
// 	return s.querier.GetNextVersion(ctx, version)
// }

// func (s *Store) GetUpdate(ctx context.Context, version string) (Update, error) {
// 	return s.querier.GetUpdate(ctx, version)
// }

// func (s *Store) GetDevice(ctx context.Context, deviceID string) (Device, error) {
// 	return s.querier.GetDevice(ctx, deviceID)
// }
