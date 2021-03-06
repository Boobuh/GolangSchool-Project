// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Boobuh/golang-school-project/dal (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dal "github.com/Boobuh/golang-school-project/dal"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateColumn mocks base method.
func (m *MockRepository) CreateColumn(arg0 *dal.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateColumn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateColumn indicates an expected call of CreateColumn.
func (mr *MockRepositoryMockRecorder) CreateColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateColumn", reflect.TypeOf((*MockRepository)(nil).CreateColumn), arg0)
}

// CreateComment mocks base method.
func (m *MockRepository) CreateComment(arg0 *dal.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockRepositoryMockRecorder) CreateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockRepository)(nil).CreateComment), arg0)
}

// CreateProject mocks base method.
func (m *MockRepository) CreateProject(arg0 *dal.Project) (*dal.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*dal.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockRepositoryMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockRepository)(nil).CreateProject), arg0)
}

// CreateTask mocks base method.
func (m *MockRepository) CreateTask(arg0 *dal.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockRepositoryMockRecorder) CreateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockRepository)(nil).CreateTask), arg0)
}

// DeleteColumn mocks base method.
func (m *MockRepository) DeleteColumn(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteColumn indicates an expected call of DeleteColumn.
func (mr *MockRepositoryMockRecorder) DeleteColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteColumn", reflect.TypeOf((*MockRepository)(nil).DeleteColumn), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockRepository) DeleteComment(arg0, arg1, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockRepositoryMockRecorder) DeleteComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockRepository)(nil).DeleteComment), arg0, arg1, arg2, arg3)
}

// DeleteProject mocks base method.
func (m *MockRepository) DeleteProject(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockRepositoryMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockRepository)(nil).DeleteProject), arg0)
}

// DeleteTask mocks base method.
func (m *MockRepository) DeleteTask(arg0, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockRepositoryMockRecorder) DeleteTask(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockRepository)(nil).DeleteTask), arg0, arg1, arg2)
}

// GetColumn mocks base method.
func (m *MockRepository) GetColumn(arg0 int) (*dal.ExtendedColumn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumn", arg0)
	ret0, _ := ret[0].(*dal.ExtendedColumn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumn indicates an expected call of GetColumn.
func (mr *MockRepositoryMockRecorder) GetColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumn", reflect.TypeOf((*MockRepository)(nil).GetColumn), arg0)
}

// GetColumns mocks base method.
func (m *MockRepository) GetColumns() ([]dal.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumns")
	ret0, _ := ret[0].([]dal.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumns indicates an expected call of GetColumns.
func (mr *MockRepositoryMockRecorder) GetColumns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumns", reflect.TypeOf((*MockRepository)(nil).GetColumns))
}

// GetComment mocks base method.
func (m *MockRepository) GetComment(arg0 int) (*dal.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0)
	ret0, _ := ret[0].(*dal.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockRepositoryMockRecorder) GetComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockRepository)(nil).GetComment), arg0)
}

// GetComments mocks base method.
func (m *MockRepository) GetComments() ([]dal.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments")
	ret0, _ := ret[0].([]dal.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockRepositoryMockRecorder) GetComments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockRepository)(nil).GetComments))
}

// GetProject mocks base method.
func (m *MockRepository) GetProject(arg0 int) (*dal.ExtendedProjectEntities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0)
	ret0, _ := ret[0].(*dal.ExtendedProjectEntities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockRepositoryMockRecorder) GetProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockRepository)(nil).GetProject), arg0)
}

// GetProjects mocks base method.
func (m *MockRepository) GetProjects() ([]dal.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects")
	ret0, _ := ret[0].([]dal.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockRepositoryMockRecorder) GetProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockRepository)(nil).GetProjects))
}

// GetTask mocks base method.
func (m *MockRepository) GetTask(arg0 int) (*dal.ExtendedTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(*dal.ExtendedTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockRepositoryMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockRepository)(nil).GetTask), arg0)
}

// GetTasks mocks base method.
func (m *MockRepository) GetTasks() ([]dal.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].([]dal.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockRepositoryMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockRepository)(nil).GetTasks))
}

// UpdateColumn mocks base method.
func (m *MockRepository) UpdateColumn(arg0 *dal.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockRepositoryMockRecorder) UpdateColumn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockRepository)(nil).UpdateColumn), arg0)
}

// UpdateComment mocks base method.
func (m *MockRepository) UpdateComment(arg0 *dal.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockRepositoryMockRecorder) UpdateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockRepository)(nil).UpdateComment), arg0)
}

// UpdateProject mocks base method.
func (m *MockRepository) UpdateProject(arg0 *dal.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockRepositoryMockRecorder) UpdateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockRepository)(nil).UpdateProject), arg0)
}

// UpdateTask mocks base method.
func (m *MockRepository) UpdateTask(arg0 *dal.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockRepositoryMockRecorder) UpdateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockRepository)(nil).UpdateTask), arg0)
}
