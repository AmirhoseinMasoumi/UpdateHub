package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/mock"
	db "github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/sqlc"
	"github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateDevice(t *testing.T) {
	device := CreateRandomdevice(t)

	testCases := []struct {
		name          string
		requestBody   CreateDeviceRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			requestBody: CreateDeviceRequest{
				DeviceId: device.DeviceID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUpdatesBetweenDates(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Update{}, nil)

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Update{
						Version:     "1.0.0",
						Description: "Initial version",
						Path:        "-",
					}, nil)

				store.EXPECT().
					CreateDevice(gomock.Any(), gomock.Eq(device.DeviceID)).
					Times(1).
					Return(device, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDevice(t, recorder.Body, device)
			},
		},
		{
			name: "NotFound",
			requestBody: CreateDeviceRequest{
				DeviceId: device.DeviceID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUpdatesBetweenDates(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Update{}, nil)

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Update{
						Version:     "1.0.0",
						Description: "Initial version",
						Path:        "-",
					}, nil)

				store.EXPECT().
					CreateDevice(gomock.Any(), gomock.Eq(device.DeviceID)).
					Times(1).
					Return(db.Device{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			requestBody: CreateDeviceRequest{
				DeviceId: device.DeviceID,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUpdatesBetweenDates(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, errors.New("internal error"))
				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateDevice(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidRequest",
			requestBody: CreateDeviceRequest{
				DeviceId: "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListUpdatesBetweenDates(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateDevice(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/device"
			requestBody, err := json.Marshal(tc.requestBody)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetAllDevices(t *testing.T) {
	devices := []db.Device{
		CreateRandomdevice(t),
		CreateRandomdevice(t),
	}
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockStore) {
				var devicesRows []db.ListAllDevicesRow
				for _, device := range devices {
					devicesRows = append(devicesRows, db.ListAllDevicesRow{
						DeviceID:      device.DeviceID,
						DeviceVersion: device.DeviceVersion,
					})
				}
				store.EXPECT().
					ListAllDevices(gomock.Any()).
					Times(1).
					Return(devicesRows, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDevices(t, recorder.Body, devices)
			},
		},
		{
			name: "NoDevices",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAllDevices(gomock.Any()).
					Times(1).
					Return([]db.ListAllDevicesRow{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAllDevices(gomock.Any()).
					Times(1).
					Return(nil, errors.New("internal error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/devices"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func CreateRandomdevice(t *testing.T) db.Device {
	return db.Device{
		DeviceID:      util.RandomString(10),
		DeviceVersion: util.RandomVersion(),
	}
}

func requireBodyMatchDevice(t *testing.T, body *bytes.Buffer, account db.Device) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getDevice db.Device
	err = json.Unmarshal(data, &getDevice)
	require.NoError(t, err)
	require.Equal(t, account, getDevice)
}

func requireBodyMatchDevices(t *testing.T, body *bytes.Buffer, accounts []db.Device) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotDevices []db.Device
	err = json.Unmarshal(data, &gotDevices)
	require.NoError(t, err)
	require.Equal(t, accounts, gotDevices)
}
