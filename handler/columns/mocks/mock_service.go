// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Boobuh/golang-school-project/handler/columns (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dal "github.com/Boobuh/golang-school-project/dal"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateColumn mocks base method.
func (m *MockService) CreateColumn(arg0 *dal.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateColumn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateColumn indicates an expected call of CreateColumn.
func (mr *MockServiceMockRecorder) CreateColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateColumn", reflect.TypeOf((*MockService)(nil).CreateColumn), arg0)
}

// DeleteColumn mocks base method.
func (m *MockService) DeleteColumn(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteColumn indicates an expected call of DeleteColumn.
func (mr *MockServiceMockRecorder) DeleteColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteColumn", reflect.TypeOf((*MockService)(nil).DeleteColumn), arg0, arg1)
}

// GetAllByProjectID mocks base method.
func (m *MockService) GetAllByProjectID(arg0 int) ([]dal.ExtendedColumn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByProjectID", arg0)
	ret0, _ := ret[0].([]dal.ExtendedColumn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByProjectID indicates an expected call of GetAllByProjectID.
func (mr *MockServiceMockRecorder) GetAllByProjectID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByProjectID", reflect.TypeOf((*MockService)(nil).GetAllByProjectID), arg0)
}

// GetColumn mocks base method.
func (m *MockService) GetColumn(arg0 int) (*dal.ExtendedColumn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumn", arg0)
	ret0, _ := ret[0].(*dal.ExtendedColumn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumn indicates an expected call of GetColumn.
func (mr *MockServiceMockRecorder) GetColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumn", reflect.TypeOf((*MockService)(nil).GetColumn), arg0)
}

// GetColumns mocks base method.
func (m *MockService) GetColumns() ([]dal.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumns")
	ret0, _ := ret[0].([]dal.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumns indicates an expected call of GetColumns.
func (mr *MockServiceMockRecorder) GetColumns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumns", reflect.TypeOf((*MockService)(nil).GetColumns))
}

// GetProjectColumn mocks base method.
func (m *MockService) GetProjectColumn(arg0, arg1 int) (*dal.ExtendedColumn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectColumn", arg0, arg1)
	ret0, _ := ret[0].(*dal.ExtendedColumn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectColumn indicates an expected call of GetProjectColumn.
func (mr *MockServiceMockRecorder) GetProjectColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectColumn", reflect.TypeOf((*MockService)(nil).GetProjectColumn), arg0, arg1)
}

// UpdateColumn mocks base method.
func (m *MockService) UpdateColumn(arg0 *dal.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockServiceMockRecorder) UpdateColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockService)(nil).UpdateColumn), arg0)
}
