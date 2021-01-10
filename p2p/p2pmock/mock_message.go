// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	p2pcommon "github.com/meeypioneer/meycoin/p2p/p2pcommon"
	types "github.com/meeypioneer/meycoin/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMessage is a mock of Message interface
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Subprotocol mocks base method
func (m *MockMessage) Subprotocol() p2pcommon.SubProtocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subprotocol")
	ret0, _ := ret[0].(p2pcommon.SubProtocol)
	return ret0
}

// Subprotocol indicates an expected call of Subprotocol
func (mr *MockMessageMockRecorder) Subprotocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subprotocol", reflect.TypeOf((*MockMessage)(nil).Subprotocol))
}

// Length mocks base method
func (m *MockMessage) Length() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockMessageMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMessage)(nil).Length))
}

// Timestamp mocks base method
func (m *MockMessage) Timestamp() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp
func (mr *MockMessageMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockMessage)(nil).Timestamp))
}

// ID mocks base method
func (m *MockMessage) ID() p2pcommon.MsgID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(p2pcommon.MsgID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockMessageMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockMessage)(nil).ID))
}

// OriginalID mocks base method
func (m *MockMessage) OriginalID() p2pcommon.MsgID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OriginalID")
	ret0, _ := ret[0].(p2pcommon.MsgID)
	return ret0
}

// OriginalID indicates an expected call of OriginalID
func (mr *MockMessageMockRecorder) OriginalID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OriginalID", reflect.TypeOf((*MockMessage)(nil).OriginalID))
}

// Payload mocks base method
func (m *MockMessage) Payload() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Payload indicates an expected call of Payload
func (mr *MockMessageMockRecorder) Payload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockMessage)(nil).Payload))
}

// MockMessageBody is a mock of MessageBody interface
type MockMessageBody struct {
	ctrl     *gomock.Controller
	recorder *MockMessageBodyMockRecorder
}

// MockMessageBodyMockRecorder is the mock recorder for MockMessageBody
type MockMessageBodyMockRecorder struct {
	mock *MockMessageBody
}

// NewMockMessageBody creates a new mock instance
func NewMockMessageBody(ctrl *gomock.Controller) *MockMessageBody {
	mock := &MockMessageBody{ctrl: ctrl}
	mock.recorder = &MockMessageBodyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageBody) EXPECT() *MockMessageBodyMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockMessageBody) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockMessageBodyMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockMessageBody)(nil).Reset))
}

// String mocks base method
func (m *MockMessageBody) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockMessageBodyMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockMessageBody)(nil).String))
}

// ProtoMessage mocks base method
func (m *MockMessageBody) ProtoMessage() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProtoMessage")
}

// ProtoMessage indicates an expected call of ProtoMessage
func (mr *MockMessageBodyMockRecorder) ProtoMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoMessage", reflect.TypeOf((*MockMessageBody)(nil).ProtoMessage))
}

// MockMessageHandler is a mock of MessageHandler interface
type MockMessageHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMessageHandlerMockRecorder
}

// MockMessageHandlerMockRecorder is the mock recorder for MockMessageHandler
type MockMessageHandlerMockRecorder struct {
	mock *MockMessageHandler
}

// NewMockMessageHandler creates a new mock instance
func NewMockMessageHandler(ctrl *gomock.Controller) *MockMessageHandler {
	mock := &MockMessageHandler{ctrl: ctrl}
	mock.recorder = &MockMessageHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageHandler) EXPECT() *MockMessageHandlerMockRecorder {
	return m.recorder
}

// AddAdvice mocks base method
func (m *MockMessageHandler) AddAdvice(advice p2pcommon.HandlerAdvice) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAdvice", advice)
}

// AddAdvice indicates an expected call of AddAdvice
func (mr *MockMessageHandlerMockRecorder) AddAdvice(advice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdvice", reflect.TypeOf((*MockMessageHandler)(nil).AddAdvice), advice)
}

// ParsePayload mocks base method
func (m *MockMessageHandler) ParsePayload(arg0 []byte) (p2pcommon.MessageBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParsePayload", arg0)
	ret0, _ := ret[0].(p2pcommon.MessageBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParsePayload indicates an expected call of ParsePayload
func (mr *MockMessageHandlerMockRecorder) ParsePayload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParsePayload", reflect.TypeOf((*MockMessageHandler)(nil).ParsePayload), arg0)
}

// CheckAuth mocks base method
func (m *MockMessageHandler) CheckAuth(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuth", msg, msgBody)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAuth indicates an expected call of CheckAuth
func (mr *MockMessageHandlerMockRecorder) CheckAuth(msg, msgBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuth", reflect.TypeOf((*MockMessageHandler)(nil).CheckAuth), msg, msgBody)
}

// Handle mocks base method
func (m *MockMessageHandler) Handle(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", msg, msgBody)
}

// Handle indicates an expected call of Handle
func (mr *MockMessageHandlerMockRecorder) Handle(msg, msgBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockMessageHandler)(nil).Handle), msg, msgBody)
}

// PreHandle mocks base method
func (m *MockMessageHandler) PreHandle() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PreHandle")
}

// PreHandle indicates an expected call of PreHandle
func (mr *MockMessageHandlerMockRecorder) PreHandle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreHandle", reflect.TypeOf((*MockMessageHandler)(nil).PreHandle))
}

// PostHandle mocks base method
func (m *MockMessageHandler) PostHandle(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostHandle", msg, msgBody)
}

// PostHandle indicates an expected call of PostHandle
func (mr *MockMessageHandlerMockRecorder) PostHandle(msg, msgBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHandle", reflect.TypeOf((*MockMessageHandler)(nil).PostHandle), msg, msgBody)
}

// MockHandlerAdvice is a mock of HandlerAdvice interface
type MockHandlerAdvice struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerAdviceMockRecorder
}

// MockHandlerAdviceMockRecorder is the mock recorder for MockHandlerAdvice
type MockHandlerAdviceMockRecorder struct {
	mock *MockHandlerAdvice
}

// NewMockHandlerAdvice creates a new mock instance
func NewMockHandlerAdvice(ctrl *gomock.Controller) *MockHandlerAdvice {
	mock := &MockHandlerAdvice{ctrl: ctrl}
	mock.recorder = &MockHandlerAdviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandlerAdvice) EXPECT() *MockHandlerAdviceMockRecorder {
	return m.recorder
}

// PreHandle mocks base method
func (m *MockHandlerAdvice) PreHandle() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PreHandle")
}

// PreHandle indicates an expected call of PreHandle
func (mr *MockHandlerAdviceMockRecorder) PreHandle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreHandle", reflect.TypeOf((*MockHandlerAdvice)(nil).PreHandle))
}

// PostHandle mocks base method
func (m *MockHandlerAdvice) PostHandle(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostHandle", msg, msgBody)
}

// PostHandle indicates an expected call of PostHandle
func (mr *MockHandlerAdviceMockRecorder) PostHandle(msg, msgBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHandle", reflect.TypeOf((*MockHandlerAdvice)(nil).PostHandle), msg, msgBody)
}

// MockAsyncHandler is a mock of AsyncHandler interface
type MockAsyncHandler struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncHandlerMockRecorder
}

// MockAsyncHandlerMockRecorder is the mock recorder for MockAsyncHandler
type MockAsyncHandlerMockRecorder struct {
	mock *MockAsyncHandler
}

// NewMockAsyncHandler creates a new mock instance
func NewMockAsyncHandler(ctrl *gomock.Controller) *MockAsyncHandler {
	mock := &MockAsyncHandler{ctrl: ctrl}
	mock.recorder = &MockAsyncHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncHandler) EXPECT() *MockAsyncHandlerMockRecorder {
	return m.recorder
}

// HandleOrNot mocks base method
func (m *MockAsyncHandler) HandleOrNot(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleOrNot", msg, msgBody)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HandleOrNot indicates an expected call of HandleOrNot
func (mr *MockAsyncHandlerMockRecorder) HandleOrNot(msg, msgBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleOrNot", reflect.TypeOf((*MockAsyncHandler)(nil).HandleOrNot), msg, msgBody)
}

// Handle mocks base method
func (m *MockAsyncHandler) Handle(msg p2pcommon.Message, msgBody p2pcommon.MessageBody, ttl time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", msg, msgBody, ttl)
}

// Handle indicates an expected call of Handle
func (mr *MockAsyncHandlerMockRecorder) Handle(msg, msgBody, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockAsyncHandler)(nil).Handle), msg, msgBody, ttl)
}

// MockMsgSigner is a mock of MsgSigner interface
type MockMsgSigner struct {
	ctrl     *gomock.Controller
	recorder *MockMsgSignerMockRecorder
}

// MockMsgSignerMockRecorder is the mock recorder for MockMsgSigner
type MockMsgSignerMockRecorder struct {
	mock *MockMsgSigner
}

// NewMockMsgSigner creates a new mock instance
func NewMockMsgSigner(ctrl *gomock.Controller) *MockMsgSigner {
	mock := &MockMsgSigner{ctrl: ctrl}
	mock.recorder = &MockMsgSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMsgSigner) EXPECT() *MockMsgSignerMockRecorder {
	return m.recorder
}

// SignMsg mocks base method
func (m *MockMsgSigner) SignMsg(msg *types.P2PMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignMsg", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignMsg indicates an expected call of SignMsg
func (mr *MockMsgSignerMockRecorder) SignMsg(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignMsg", reflect.TypeOf((*MockMsgSigner)(nil).SignMsg), msg)
}

// VerifyMsg mocks base method
func (m *MockMsgSigner) VerifyMsg(msg *types.P2PMessage, senderID types.PeerID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMsg", msg, senderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyMsg indicates an expected call of VerifyMsg
func (mr *MockMsgSignerMockRecorder) VerifyMsg(msg, senderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMsg", reflect.TypeOf((*MockMsgSigner)(nil).VerifyMsg), msg, senderID)
}