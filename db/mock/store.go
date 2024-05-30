// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockStore) CreateDevice(arg0 context.Context, arg1 string) (db.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0, arg1)
	ret0, _ := ret[0].(db.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockStoreMockRecorder) CreateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockStore)(nil).CreateDevice), arg0, arg1)
}

// CreateUpdate mocks base method.
func (m *MockStore) CreateUpdate(arg0 context.Context, arg1 db.CreateUpdateParams) (db.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpdate indicates an expected call of CreateUpdate.
func (mr *MockStoreMockRecorder) CreateUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpdate", reflect.TypeOf((*MockStore)(nil).CreateUpdate), arg0, arg1)
}

// DeleteDevice mocks base method.
func (m *MockStore) DeleteDevice(arg0 context.Context, arg1 string) (db.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", arg0, arg1)
	ret0, _ := ret[0].(db.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockStoreMockRecorder) DeleteDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockStore)(nil).DeleteDevice), arg0, arg1)
}

// DeleteUpdate mocks base method.
func (m *MockStore) DeleteUpdate(arg0 context.Context, arg1 string) (db.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUpdate indicates an expected call of DeleteUpdate.
func (mr *MockStoreMockRecorder) DeleteUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUpdate", reflect.TypeOf((*MockStore)(nil).DeleteUpdate), arg0, arg1)
}

// GetDevice mocks base method.
func (m *MockStore) GetDevice(arg0 context.Context, arg1 string) (db.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", arg0, arg1)
	ret0, _ := ret[0].(db.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockStoreMockRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockStore)(nil).GetDevice), arg0, arg1)
}

// GetNextVersion mocks base method.
func (m *MockStore) GetNextVersion(arg0 context.Context, arg1 string) (db.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextVersion", arg0, arg1)
	ret0, _ := ret[0].(db.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextVersion indicates an expected call of GetNextVersion.
func (mr *MockStoreMockRecorder) GetNextVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextVersion", reflect.TypeOf((*MockStore)(nil).GetNextVersion), arg0, arg1)
}

// GetUpdate mocks base method.
func (m *MockStore) GetUpdate(arg0 context.Context, arg1 string) (db.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdate indicates an expected call of GetUpdate.
func (mr *MockStoreMockRecorder) GetUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdate", reflect.TypeOf((*MockStore)(nil).GetUpdate), arg0, arg1)
}

// ListAllDevices mocks base method.
func (m *MockStore) ListAllDevices(arg0 context.Context) ([]db.ListAllDevicesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDevices", arg0)
	ret0, _ := ret[0].([]db.ListAllDevicesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDevices indicates an expected call of ListAllDevices.
func (mr *MockStoreMockRecorder) ListAllDevices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDevices", reflect.TypeOf((*MockStore)(nil).ListAllDevices), arg0)
}

// ListUpdatesBetweenDates mocks base method.
func (m *MockStore) ListUpdatesBetweenDates(arg0 context.Context, arg1 db.ListUpdatesBetweenDatesParams) ([]db.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUpdatesBetweenDates", arg0, arg1)
	ret0, _ := ret[0].([]db.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUpdatesBetweenDates indicates an expected call of ListUpdatesBetweenDates.
func (mr *MockStoreMockRecorder) ListUpdatesBetweenDates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpdatesBetweenDates", reflect.TypeOf((*MockStore)(nil).ListUpdatesBetweenDates), arg0, arg1)
}

// UpdateDevice mocks base method.
func (m *MockStore) UpdateDevice(arg0 context.Context, arg1 db.UpdateDeviceParams) (db.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0, arg1)
	ret0, _ := ret[0].(db.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockStoreMockRecorder) UpdateDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockStore)(nil).UpdateDevice), arg0, arg1)
}
