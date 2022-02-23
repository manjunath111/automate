// Code generated by MockGen. DO NOT EDIT.
// Source: infra_proxy/migrations/migrations.pb.go

// Package migrations is a generated GoMock package.
package migrations

import (
	context "context"
	request "github.com/chef/automate/api/external/infra_proxy/migrations/request"
	response "github.com/chef/automate/api/external/infra_proxy/migrations/response"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockInfraProxyMigrationServiceClient is a mock of InfraProxyMigrationServiceClient interface
type MockInfraProxyMigrationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInfraProxyMigrationServiceClientMockRecorder
}

// MockInfraProxyMigrationServiceClientMockRecorder is the mock recorder for MockInfraProxyMigrationServiceClient
type MockInfraProxyMigrationServiceClientMockRecorder struct {
	mock *MockInfraProxyMigrationServiceClient
}

// NewMockInfraProxyMigrationServiceClient creates a new mock instance
func NewMockInfraProxyMigrationServiceClient(ctrl *gomock.Controller) *MockInfraProxyMigrationServiceClient {
	mock := &MockInfraProxyMigrationServiceClient{ctrl: ctrl}
	mock.recorder = &MockInfraProxyMigrationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfraProxyMigrationServiceClient) EXPECT() *MockInfraProxyMigrationServiceClientMockRecorder {
	return m.recorder
}

// GetMigrationStatus mocks base method
func (m *MockInfraProxyMigrationServiceClient) GetMigrationStatus(ctx context.Context, in *request.GetMigrationStatusRequest, opts ...grpc.CallOption) (*response.GetMigrationStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMigrationStatus", varargs...)
	ret0, _ := ret[0].(*response.GetMigrationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigrationStatus indicates an expected call of GetMigrationStatus
func (mr *MockInfraProxyMigrationServiceClientMockRecorder) GetMigrationStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrationStatus", reflect.TypeOf((*MockInfraProxyMigrationServiceClient)(nil).GetMigrationStatus), varargs...)
}

// CancelMigration mocks base method
func (m *MockInfraProxyMigrationServiceClient) CancelMigration(ctx context.Context, in *request.CancelMigrationRequest, opts ...grpc.CallOption) (*response.CancelMigrationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelMigration", varargs...)
	ret0, _ := ret[0].(*response.CancelMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelMigration indicates an expected call of CancelMigration
func (mr *MockInfraProxyMigrationServiceClientMockRecorder) CancelMigration(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelMigration", reflect.TypeOf((*MockInfraProxyMigrationServiceClient)(nil).CancelMigration), varargs...)
}

// GetStagedData mocks base method
func (m *MockInfraProxyMigrationServiceClient) GetStagedData(ctx context.Context, in *request.GetStagedDataRequest, opts ...grpc.CallOption) (*response.GetStagedDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStagedData", varargs...)
	ret0, _ := ret[0].(*response.GetStagedDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStagedData indicates an expected call of GetStagedData
func (mr *MockInfraProxyMigrationServiceClientMockRecorder) GetStagedData(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStagedData", reflect.TypeOf((*MockInfraProxyMigrationServiceClient)(nil).GetStagedData), varargs...)
}

// ConfirmPreview mocks base method
func (m *MockInfraProxyMigrationServiceClient) ConfirmPreview(ctx context.Context, in *request.ConfirmPreview, opts ...grpc.CallOption) (*response.ConfirmPreview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfirmPreview", varargs...)
	ret0, _ := ret[0].(*response.ConfirmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmPreview indicates an expected call of ConfirmPreview
func (mr *MockInfraProxyMigrationServiceClientMockRecorder) ConfirmPreview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmPreview", reflect.TypeOf((*MockInfraProxyMigrationServiceClient)(nil).ConfirmPreview), varargs...)
}

// MockInfraProxyMigrationServiceServer is a mock of InfraProxyMigrationServiceServer interface
type MockInfraProxyMigrationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockInfraProxyMigrationServiceServerMockRecorder
}

// MockInfraProxyMigrationServiceServerMockRecorder is the mock recorder for MockInfraProxyMigrationServiceServer
type MockInfraProxyMigrationServiceServerMockRecorder struct {
	mock *MockInfraProxyMigrationServiceServer
}

// NewMockInfraProxyMigrationServiceServer creates a new mock instance
func NewMockInfraProxyMigrationServiceServer(ctrl *gomock.Controller) *MockInfraProxyMigrationServiceServer {
	mock := &MockInfraProxyMigrationServiceServer{ctrl: ctrl}
	mock.recorder = &MockInfraProxyMigrationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfraProxyMigrationServiceServer) EXPECT() *MockInfraProxyMigrationServiceServerMockRecorder {
	return m.recorder
}

// GetMigrationStatus mocks base method
func (m *MockInfraProxyMigrationServiceServer) GetMigrationStatus(arg0 context.Context, arg1 *request.GetMigrationStatusRequest) (*response.GetMigrationStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrationStatus", arg0, arg1)
	ret0, _ := ret[0].(*response.GetMigrationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigrationStatus indicates an expected call of GetMigrationStatus
func (mr *MockInfraProxyMigrationServiceServerMockRecorder) GetMigrationStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrationStatus", reflect.TypeOf((*MockInfraProxyMigrationServiceServer)(nil).GetMigrationStatus), arg0, arg1)
}

// CancelMigration mocks base method
func (m *MockInfraProxyMigrationServiceServer) CancelMigration(arg0 context.Context, arg1 *request.CancelMigrationRequest) (*response.CancelMigrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelMigration", arg0, arg1)
	ret0, _ := ret[0].(*response.CancelMigrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelMigration indicates an expected call of CancelMigration
func (mr *MockInfraProxyMigrationServiceServerMockRecorder) CancelMigration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelMigration", reflect.TypeOf((*MockInfraProxyMigrationServiceServer)(nil).CancelMigration), arg0, arg1)
}

// GetStagedData mocks base method
func (m *MockInfraProxyMigrationServiceServer) GetStagedData(arg0 context.Context, arg1 *request.GetStagedDataRequest) (*response.GetStagedDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStagedData", arg0, arg1)
	ret0, _ := ret[0].(*response.GetStagedDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStagedData indicates an expected call of GetStagedData
func (mr *MockInfraProxyMigrationServiceServerMockRecorder) GetStagedData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStagedData", reflect.TypeOf((*MockInfraProxyMigrationServiceServer)(nil).GetStagedData), arg0, arg1)
}

// ConfirmPreview mocks base method
func (m *MockInfraProxyMigrationServiceServer) ConfirmPreview(arg0 context.Context, arg1 *request.ConfirmPreview) (*response.ConfirmPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmPreview", arg0, arg1)
	ret0, _ := ret[0].(*response.ConfirmPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmPreview indicates an expected call of ConfirmPreview
func (mr *MockInfraProxyMigrationServiceServerMockRecorder) ConfirmPreview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmPreview", reflect.TypeOf((*MockInfraProxyMigrationServiceServer)(nil).ConfirmPreview), arg0, arg1)
}