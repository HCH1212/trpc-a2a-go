// Code generated by MockGen. DO NOT EDIT.
// Source: stub/chat/protobuf/chat.trpc.go
//
// Generated by this command:
//
//	mockgen -destination=stub/chat/protobuf/chat_mock.go -package=chat -self_package=chat/protobuf --source=stub/chat/protobuf/chat.trpc.go
//

// Package chat is a generated GoMock package.
package chat

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockChatroomServiceService is a mock of ChatroomServiceService interface.
type MockChatroomServiceService struct {
	ctrl     *gomock.Controller
	recorder *MockChatroomServiceServiceMockRecorder
}

// MockChatroomServiceServiceMockRecorder is the mock recorder for MockChatroomServiceService.
type MockChatroomServiceServiceMockRecorder struct {
	mock *MockChatroomServiceService
}

// NewMockChatroomServiceService creates a new mock instance.
func NewMockChatroomServiceService(ctrl *gomock.Controller) *MockChatroomServiceService {
	mock := &MockChatroomServiceService{ctrl: ctrl}
	mock.recorder = &MockChatroomServiceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatroomServiceService) EXPECT() *MockChatroomServiceServiceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockChatroomServiceService) ISGOMOCK() struct{} {
	return struct{}{}
}

// Chat mocks base method.
func (m *MockChatroomServiceService) Chat(arg0 ChatroomService_ChatServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chat", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chat indicates an expected call of Chat.
func (mr *MockChatroomServiceServiceMockRecorder) Chat(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatroomServiceService)(nil).Chat), arg0)
}

// MockChatroomService_ChatServer is a mock of ChatroomService_ChatServer interface.
type MockChatroomService_ChatServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatroomService_ChatServerMockRecorder
}

// MockChatroomService_ChatServerMockRecorder is the mock recorder for MockChatroomService_ChatServer.
type MockChatroomService_ChatServerMockRecorder struct {
	mock *MockChatroomService_ChatServer
}

// NewMockChatroomService_ChatServer creates a new mock instance.
func NewMockChatroomService_ChatServer(ctrl *gomock.Controller) *MockChatroomService_ChatServer {
	mock := &MockChatroomService_ChatServer{ctrl: ctrl}
	mock.recorder = &MockChatroomService_ChatServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatroomService_ChatServer) EXPECT() *MockChatroomService_ChatServerMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockChatroomService_ChatServer) ISGOMOCK() struct{} {
	return struct{}{}
}

// Context mocks base method.
func (m *MockChatroomService_ChatServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatroomService_ChatServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatroomService_ChatServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockChatroomService_ChatServer) Recv() (*ChatroomMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ChatroomMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockChatroomService_ChatServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChatroomService_ChatServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockChatroomService_ChatServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatroomService_ChatServerMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatroomService_ChatServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockChatroomService_ChatServer) Send(arg0 *ChatroomMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockChatroomService_ChatServerMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatroomService_ChatServer)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockChatroomService_ChatServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatroomService_ChatServerMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatroomService_ChatServer)(nil).SendMsg), m)
}

// MockChatroomServiceClientProxy is a mock of ChatroomServiceClientProxy interface.
type MockChatroomServiceClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockChatroomServiceClientProxyMockRecorder
}

// MockChatroomServiceClientProxyMockRecorder is the mock recorder for MockChatroomServiceClientProxy.
type MockChatroomServiceClientProxyMockRecorder struct {
	mock *MockChatroomServiceClientProxy
}

// NewMockChatroomServiceClientProxy creates a new mock instance.
func NewMockChatroomServiceClientProxy(ctrl *gomock.Controller) *MockChatroomServiceClientProxy {
	mock := &MockChatroomServiceClientProxy{ctrl: ctrl}
	mock.recorder = &MockChatroomServiceClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatroomServiceClientProxy) EXPECT() *MockChatroomServiceClientProxyMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockChatroomServiceClientProxy) ISGOMOCK() struct{} {
	return struct{}{}
}

// Chat mocks base method.
func (m *MockChatroomServiceClientProxy) Chat(ctx context.Context, opts ...client.Option) (ChatroomService_ChatClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Chat", varargs...)
	ret0, _ := ret[0].(ChatroomService_ChatClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat.
func (mr *MockChatroomServiceClientProxyMockRecorder) Chat(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatroomServiceClientProxy)(nil).Chat), varargs...)
}

// MockChatroomService_ChatClient is a mock of ChatroomService_ChatClient interface.
type MockChatroomService_ChatClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatroomService_ChatClientMockRecorder
}

// MockChatroomService_ChatClientMockRecorder is the mock recorder for MockChatroomService_ChatClient.
type MockChatroomService_ChatClientMockRecorder struct {
	mock *MockChatroomService_ChatClient
}

// NewMockChatroomService_ChatClient creates a new mock instance.
func NewMockChatroomService_ChatClient(ctrl *gomock.Controller) *MockChatroomService_ChatClient {
	mock := &MockChatroomService_ChatClient{ctrl: ctrl}
	mock.recorder = &MockChatroomService_ChatClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatroomService_ChatClient) EXPECT() *MockChatroomService_ChatClientMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockChatroomService_ChatClient) ISGOMOCK() struct{} {
	return struct{}{}
}

// CloseSend mocks base method.
func (m *MockChatroomService_ChatClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockChatroomService_ChatClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockChatroomService_ChatClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatroomService_ChatClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).Context))
}

// Recv mocks base method.
func (m *MockChatroomService_ChatClient) Recv() (*ChatroomMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ChatroomMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockChatroomService_ChatClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockChatroomService_ChatClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatroomService_ChatClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockChatroomService_ChatClient) Send(arg0 *ChatroomMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockChatroomService_ChatClientMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockChatroomService_ChatClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatroomService_ChatClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatroomService_ChatClient)(nil).SendMsg), m)
}
