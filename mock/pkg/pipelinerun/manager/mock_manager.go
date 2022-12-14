// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	reflect "reflect"

	q "g.hz.netease.com/horizon/lib/q"
	models "g.hz.netease.com/horizon/pkg/pipelinerun/models"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockManager) Create(ctx context.Context, pipelinerun *models.Pipelinerun) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, pipelinerun)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockManagerMockRecorder) Create(ctx, pipelinerun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, pipelinerun)
}

// DeleteByClusterID mocks base method.
func (m *MockManager) DeleteByClusterID(ctx context.Context, clusterID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByClusterID", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByClusterID indicates an expected call of DeleteByClusterID.
func (mr *MockManagerMockRecorder) DeleteByClusterID(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByClusterID", reflect.TypeOf((*MockManager)(nil).DeleteByClusterID), ctx, clusterID)
}

// DeleteByID mocks base method.
func (m *MockManager) DeleteByID(ctx context.Context, pipelinerunID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, pipelinerunID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockManagerMockRecorder) DeleteByID(ctx, pipelinerunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockManager)(nil).DeleteByID), ctx, pipelinerunID)
}

// GetByCIEventID mocks base method.
func (m *MockManager) GetByCIEventID(ctx context.Context, ciEventID string) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCIEventID", ctx, ciEventID)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCIEventID indicates an expected call of GetByCIEventID.
func (mr *MockManagerMockRecorder) GetByCIEventID(ctx, ciEventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCIEventID", reflect.TypeOf((*MockManager)(nil).GetByCIEventID), ctx, ciEventID)
}

// GetByClusterID mocks base method.
func (m *MockManager) GetByClusterID(ctx context.Context, clusterID uint, canRollback bool, query q.Query) (int, []*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByClusterID", ctx, clusterID, canRollback, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]*models.Pipelinerun)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByClusterID indicates an expected call of GetByClusterID.
func (mr *MockManagerMockRecorder) GetByClusterID(ctx, clusterID, canRollback, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByClusterID", reflect.TypeOf((*MockManager)(nil).GetByClusterID), ctx, clusterID, canRollback, query)
}

// GetByID mocks base method.
func (m *MockManager) GetByID(ctx context.Context, pipelinerunID uint) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, pipelinerunID)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockManagerMockRecorder) GetByID(ctx, pipelinerunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockManager)(nil).GetByID), ctx, pipelinerunID)
}

// GetFirstCanRollbackPipelinerun mocks base method.
func (m *MockManager) GetFirstCanRollbackPipelinerun(ctx context.Context, clusterID uint) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstCanRollbackPipelinerun", ctx, clusterID)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstCanRollbackPipelinerun indicates an expected call of GetFirstCanRollbackPipelinerun.
func (mr *MockManagerMockRecorder) GetFirstCanRollbackPipelinerun(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstCanRollbackPipelinerun", reflect.TypeOf((*MockManager)(nil).GetFirstCanRollbackPipelinerun), ctx, clusterID)
}

// GetLatestByClusterIDAndAction mocks base method.
func (m *MockManager) GetLatestByClusterIDAndAction(ctx context.Context, clusterID uint, action string) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestByClusterIDAndAction", ctx, clusterID, action)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestByClusterIDAndAction indicates an expected call of GetLatestByClusterIDAndAction.
func (mr *MockManagerMockRecorder) GetLatestByClusterIDAndAction(ctx, clusterID, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestByClusterIDAndAction", reflect.TypeOf((*MockManager)(nil).GetLatestByClusterIDAndAction), ctx, clusterID, action)
}

// GetLatestByClusterIDAndActionAndStatus mocks base method.
func (m *MockManager) GetLatestByClusterIDAndActionAndStatus(ctx context.Context, clusterID uint, action, status string) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestByClusterIDAndActionAndStatus", ctx, clusterID, action, status)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestByClusterIDAndActionAndStatus indicates an expected call of GetLatestByClusterIDAndActionAndStatus.
func (mr *MockManagerMockRecorder) GetLatestByClusterIDAndActionAndStatus(ctx, clusterID, action, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestByClusterIDAndActionAndStatus", reflect.TypeOf((*MockManager)(nil).GetLatestByClusterIDAndActionAndStatus), ctx, clusterID, action, status)
}

// GetLatestSuccessByClusterID mocks base method.
func (m *MockManager) GetLatestSuccessByClusterID(ctx context.Context, clusterID uint) (*models.Pipelinerun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestSuccessByClusterID", ctx, clusterID)
	ret0, _ := ret[0].(*models.Pipelinerun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestSuccessByClusterID indicates an expected call of GetLatestSuccessByClusterID.
func (mr *MockManagerMockRecorder) GetLatestSuccessByClusterID(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestSuccessByClusterID", reflect.TypeOf((*MockManager)(nil).GetLatestSuccessByClusterID), ctx, clusterID)
}

// UpdateCIEventIDByID mocks base method.
func (m *MockManager) UpdateCIEventIDByID(ctx context.Context, pipelinerunID uint, ciEventID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCIEventIDByID", ctx, pipelinerunID, ciEventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCIEventIDByID indicates an expected call of UpdateCIEventIDByID.
func (mr *MockManagerMockRecorder) UpdateCIEventIDByID(ctx, pipelinerunID, ciEventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCIEventIDByID", reflect.TypeOf((*MockManager)(nil).UpdateCIEventIDByID), ctx, pipelinerunID, ciEventID)
}

// UpdateConfigCommitByID mocks base method.
func (m *MockManager) UpdateConfigCommitByID(ctx context.Context, pipelinerunID uint, commit string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigCommitByID", ctx, pipelinerunID, commit)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigCommitByID indicates an expected call of UpdateConfigCommitByID.
func (mr *MockManagerMockRecorder) UpdateConfigCommitByID(ctx, pipelinerunID, commit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigCommitByID", reflect.TypeOf((*MockManager)(nil).UpdateConfigCommitByID), ctx, pipelinerunID, commit)
}

// UpdateResultByID mocks base method.
func (m *MockManager) UpdateResultByID(ctx context.Context, pipelinerunID uint, result *models.Result) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResultByID", ctx, pipelinerunID, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResultByID indicates an expected call of UpdateResultByID.
func (mr *MockManagerMockRecorder) UpdateResultByID(ctx, pipelinerunID, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResultByID", reflect.TypeOf((*MockManager)(nil).UpdateResultByID), ctx, pipelinerunID, result)
}

// UpdateStatusByID mocks base method.
func (m *MockManager) UpdateStatusByID(ctx context.Context, pipelinerunID uint, result models.PipelineStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusByID", ctx, pipelinerunID, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusByID indicates an expected call of UpdateStatusByID.
func (mr *MockManagerMockRecorder) UpdateStatusByID(ctx, pipelinerunID, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusByID", reflect.TypeOf((*MockManager)(nil).UpdateStatusByID), ctx, pipelinerunID, result)
}
