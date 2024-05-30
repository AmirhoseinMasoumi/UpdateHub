// updates_test.go
package db

import (
	"context"
	"testing"
	"time"

	"github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func CreateVersionUpdate(t *testing.T, version string) Update {
	arg := CreateUpdateParams{
		Version:     version,
		Path:        util.RandomPath(),
		Description: util.RandomString(20),
	}

	update, err := testQueries.CreateUpdate(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Version, update.Version)
	require.Equal(t, arg.Path, update.Path)
	require.Equal(t, arg.Description, update.Description)
	return update
}

func CreateRandomUpdate(t *testing.T) Update {
	arg := CreateUpdateParams{
		Version:     util.RandomVersion(),
		Path:        util.RandomPath(),
		Description: util.RandomString(20),
		Checksum:    util.RandomString(10),
	}

	update, err := testQueries.CreateUpdate(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Version, update.Version)
	require.Equal(t, arg.Path, update.Path)
	require.Equal(t, arg.Description, update.Description)
	require.Equal(t, arg.Checksum, update.Checksum)

	return update
}

func TestCreateUpdate(t *testing.T) {
	CreateRandomUpdate(t)
}

func TestCheckUpdate(t *testing.T) {
	update1 := CreateVersionUpdate(t, "1.0.0")
	update2 := CreateVersionUpdate(t, "1.0.1")

	update, err := testQueries.GetNextVersion(context.Background(), update1.Version)
	require.NoError(t, err)
	require.Equal(t, update.Version, update2.Version)
}

func TestDeleteUpdate(t *testing.T) {
	update := CreateRandomUpdate(t)

	result, err := testQueries.DeleteUpdate(context.Background(), update.Version)
	require.NoError(t, err)
	require.Equal(t, update.Version, result.Version)
}

func TestListUpdates(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUpdate(t)
	}

	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()

	arg := ListUpdatesBetweenDatesParams{
		StartDate: yesterday,
		EndDate:   today,
	}

	updates, err := testQueries.ListUpdatesBetweenDates(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updates)

	for _, account := range updates {
		require.NotEmpty(t, account)
	}
}

func TestGetNextVersion(t *testing.T) {
	update1 := CreateVersionUpdate(t, "2.0.0")
	update2 := CreateVersionUpdate(t, "2.0.1")

	nextUpdate, err := testQueries.GetNextVersion(context.Background(), update1.Version)
	require.NoError(t, err)
	require.NotEmpty(t, nextUpdate)
	require.Equal(t, update2.Version, nextUpdate.Version)

}

func TestGetUpdate(t *testing.T) {
	update := CreateRandomUpdate(t)

	result, err := testQueries.GetUpdate(context.Background(), update.Version)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, update.Version, result.Version)
	require.Equal(t, update.Path, result.Path)
	require.Equal(t, update.Description, result.Description)
	require.Equal(t, update.Checksum, result.Checksum)
}
