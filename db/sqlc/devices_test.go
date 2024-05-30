package db

import (
	"context"
	"testing"
	"time"

	"github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func CreateRandomDevice(t *testing.T) Device {
	update := CreateVersionUpdate(t, "1.0.0")

	expectedDevice := Device{
		DeviceID:      util.RandomString(10),
		DeviceVersion: update.Version,
	}

	device, err := testQueries.CreateDevice(context.Background(), expectedDevice.DeviceID)
	require.NoError(t, err)
	require.Equal(t, expectedDevice.DeviceID, device.DeviceID)
	require.Equal(t, expectedDevice.DeviceVersion, device.DeviceVersion)

	return device
}

func TestCreateDevice(t *testing.T) {
	CreateRandomDevice(t)
}

func TestDeleteDevice(t *testing.T) {
	device := CreateRandomDevice(t)

	result, err := testQueries.DeleteDevice(context.Background(), device.DeviceID)
	require.NoError(t, err)
	require.Equal(t, device.DeviceID, result.DeviceID)
}

func TestUpdateDevice(t *testing.T) {
	device := CreateRandomDevice(t)
	update := CreateRandomUpdate(t)

	arg := UpdateDeviceParams{
		DeviceID:      device.DeviceID,
		DeviceVersion: update.Version,
		LastUpdate:    time.Now(),
	}

	device, err := testQueries.UpdateDevice(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.DeviceID, device.DeviceID)
	require.Equal(t, arg.DeviceVersion, device.DeviceVersion)
}

func TestListDevices(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomDevice(t)
	}

	devices, err := testQueries.ListAllDevices(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, devices)

	for _, device := range devices {
		require.NotEmpty(t, device)
	}
}

func TestGetDevice(t *testing.T) {
	update := CreateRandomDevice(t)

	result, err := testQueries.GetDevice(context.Background(), update.DeviceID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, update.DeviceID, result.DeviceID)
}
