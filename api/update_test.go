package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/mock"
	db "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateUpdate(t *testing.T) {
	update := db.Update{
		Version:     "1.0.0",
		Description: "Test Update",
		Path:        "/test/update.zip",
	}

	testCases := []struct {
		name          string
		requestBody   CreateUpdateRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ValidRequest",
			requestBody: CreateUpdateRequest{
				Version:     "1.0.0",
				Description: "Test Update",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Update{}, sql.ErrNoRows)

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(update, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUpdate(t, recorder.Body, update)
			},
		},
		{
			name: "BadRequest",
			requestBody: CreateUpdateRequest{
				Version:     "", // Missing version
				Description: "Test Update",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUpdate(gomock.Any(), gomock.Any()).
					Times(0)

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "DuplicateUpdate",
			requestBody: CreateUpdateRequest{
				Version:     "1.0.0",
				Description: "Test Update",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(update, nil)

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			requestBody: CreateUpdateRequest{
				Version:     "1.0.0",
				Description: "Test Update",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUpdate(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Update{}, errors.New("internal error"))

				store.EXPECT().
					CreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)
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

			url := "/update"
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

func requireBodyMatchUpdate(t *testing.T, body *bytes.Buffer, update db.Update) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUpdate db.Update
	err = json.Unmarshal(data, &gotUpdate)
	require.NoError(t, err)
	require.Equal(t, update, gotUpdate)
}
