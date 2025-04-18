// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chef/automate/api/interservice/compliance/ingest/ingest (interfaces: ComplianceIngesterServiceClient,ComplianceIngesterService_ProcessComplianceReportClient,ComplianceIngesterServiceServer,ComplianceIngesterService_ProcessComplianceReportServer)

// Package ingest is a generated GoMock package.
package ingest

import (
	context "context"
	reflect "reflect"

	event "github.com/chef/automate/api/interservice/event"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockComplianceIngesterServiceClient is a mock of ComplianceIngesterServiceClient interface.
type MockComplianceIngesterServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceIngesterServiceClientMockRecorder
}

// MockComplianceIngesterServiceClientMockRecorder is the mock recorder for MockComplianceIngesterServiceClient.
type MockComplianceIngesterServiceClientMockRecorder struct {
	mock *MockComplianceIngesterServiceClient
}

// NewMockComplianceIngesterServiceClient creates a new mock instance.
func NewMockComplianceIngesterServiceClient(ctrl *gomock.Controller) *MockComplianceIngesterServiceClient {
	mock := &MockComplianceIngesterServiceClient{ctrl: ctrl}
	mock.recorder = &MockComplianceIngesterServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceIngesterServiceClient) EXPECT() *MockComplianceIngesterServiceClientMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method.
func (m *MockComplianceIngesterServiceClient) HandleEvent(arg0 context.Context, arg1 *event.EventMsg, arg2 ...grpc.CallOption) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandleEvent", varargs...)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent.
func (mr *MockComplianceIngesterServiceClientMockRecorder) HandleEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockComplianceIngesterServiceClient)(nil).HandleEvent), varargs...)
}

// ProcessComplianceReport mocks base method.
func (m *MockComplianceIngesterServiceClient) ProcessComplianceReport(arg0 context.Context, arg1 ...grpc.CallOption) (ComplianceIngesterService_ProcessComplianceReportClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProcessComplianceReport", varargs...)
	ret0, _ := ret[0].(ComplianceIngesterService_ProcessComplianceReportClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessComplianceReport indicates an expected call of ProcessComplianceReport.
func (mr *MockComplianceIngesterServiceClientMockRecorder) ProcessComplianceReport(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessComplianceReport", reflect.TypeOf((*MockComplianceIngesterServiceClient)(nil).ProcessComplianceReport), varargs...)
}

// ProjectUpdateStatus mocks base method.
func (m *MockComplianceIngesterServiceClient) ProjectUpdateStatus(arg0 context.Context, arg1 *ProjectUpdateStatusReq, arg2 ...grpc.CallOption) (*ProjectUpdateStatusResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProjectUpdateStatus", varargs...)
	ret0, _ := ret[0].(*ProjectUpdateStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateStatus indicates an expected call of ProjectUpdateStatus.
func (mr *MockComplianceIngesterServiceClientMockRecorder) ProjectUpdateStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateStatus", reflect.TypeOf((*MockComplianceIngesterServiceClient)(nil).ProjectUpdateStatus), varargs...)
}

// MockComplianceIngesterService_ProcessComplianceReportClient is a mock of ComplianceIngesterService_ProcessComplianceReportClient interface.
type MockComplianceIngesterService_ProcessComplianceReportClient struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder
}

// MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder is the mock recorder for MockComplianceIngesterService_ProcessComplianceReportClient.
type MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder struct {
	mock *MockComplianceIngesterService_ProcessComplianceReportClient
}

// NewMockComplianceIngesterService_ProcessComplianceReportClient creates a new mock instance.
func NewMockComplianceIngesterService_ProcessComplianceReportClient(ctrl *gomock.Controller) *MockComplianceIngesterService_ProcessComplianceReportClient {
	mock := &MockComplianceIngesterService_ProcessComplianceReportClient{ctrl: ctrl}
	mock.recorder = &MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) EXPECT() *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) CloseAndRecv() (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).Context))
}

// Header mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) Send(arg0 *ReportData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockComplianceIngesterService_ProcessComplianceReportClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportClient)(nil).Trailer))
}

// MockComplianceIngesterServiceServer is a mock of ComplianceIngesterServiceServer interface.
type MockComplianceIngesterServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceIngesterServiceServerMockRecorder
}

// MockComplianceIngesterServiceServerMockRecorder is the mock recorder for MockComplianceIngesterServiceServer.
type MockComplianceIngesterServiceServerMockRecorder struct {
	mock *MockComplianceIngesterServiceServer
}

// NewMockComplianceIngesterServiceServer creates a new mock instance.
func NewMockComplianceIngesterServiceServer(ctrl *gomock.Controller) *MockComplianceIngesterServiceServer {
	mock := &MockComplianceIngesterServiceServer{ctrl: ctrl}
	mock.recorder = &MockComplianceIngesterServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceIngesterServiceServer) EXPECT() *MockComplianceIngesterServiceServerMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method.
func (m *MockComplianceIngesterServiceServer) HandleEvent(arg0 context.Context, arg1 *event.EventMsg) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvent", arg0, arg1)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent.
func (mr *MockComplianceIngesterServiceServerMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockComplianceIngesterServiceServer)(nil).HandleEvent), arg0, arg1)
}

// ProcessComplianceReport mocks base method.
func (m *MockComplianceIngesterServiceServer) ProcessComplianceReport(arg0 ComplianceIngesterService_ProcessComplianceReportServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessComplianceReport", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessComplianceReport indicates an expected call of ProcessComplianceReport.
func (mr *MockComplianceIngesterServiceServerMockRecorder) ProcessComplianceReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessComplianceReport", reflect.TypeOf((*MockComplianceIngesterServiceServer)(nil).ProcessComplianceReport), arg0)
}

// ProjectUpdateStatus mocks base method.
func (m *MockComplianceIngesterServiceServer) ProjectUpdateStatus(arg0 context.Context, arg1 *ProjectUpdateStatusReq) (*ProjectUpdateStatusResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdateStatus", arg0, arg1)
	ret0, _ := ret[0].(*ProjectUpdateStatusResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdateStatus indicates an expected call of ProjectUpdateStatus.
func (mr *MockComplianceIngesterServiceServerMockRecorder) ProjectUpdateStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdateStatus", reflect.TypeOf((*MockComplianceIngesterServiceServer)(nil).ProjectUpdateStatus), arg0, arg1)
}

// MockComplianceIngesterService_ProcessComplianceReportServer is a mock of ComplianceIngesterService_ProcessComplianceReportServer interface.
type MockComplianceIngesterService_ProcessComplianceReportServer struct {
	ctrl     *gomock.Controller
	recorder *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder
}

// MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder is the mock recorder for MockComplianceIngesterService_ProcessComplianceReportServer.
type MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder struct {
	mock *MockComplianceIngesterService_ProcessComplianceReportServer
}

// NewMockComplianceIngesterService_ProcessComplianceReportServer creates a new mock instance.
func NewMockComplianceIngesterService_ProcessComplianceReportServer(ctrl *gomock.Controller) *MockComplianceIngesterService_ProcessComplianceReportServer {
	mock := &MockComplianceIngesterService_ProcessComplianceReportServer{ctrl: ctrl}
	mock.recorder = &MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) EXPECT() *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) Recv() (*ReportData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ReportData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).RecvMsg), arg0)
}

// SendAndClose mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) SendAndClose(arg0 *emptypb.Empty) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockComplianceIngesterService_ProcessComplianceReportServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockComplianceIngesterService_ProcessComplianceReportServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockComplianceIngesterService_ProcessComplianceReportServer)(nil).SetTrailer), arg0)
}
