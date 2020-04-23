// Code generated by MockGen. DO NOT EDIT.
// Source: hostv2.go

// Package mock_hostv2 is a generated GoMock package.
package mock_hostv2

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	reflect "reflect"
)

// MocktopicHandle is a mock of topicHandle interface
type MocktopicHandle struct {
	ctrl     *gomock.Controller
	recorder *MocktopicHandleMockRecorder
}

// MocktopicHandleMockRecorder is the mock recorder for MocktopicHandle
type MocktopicHandleMockRecorder struct {
	mock *MocktopicHandle
}

// NewMocktopicHandle creates a new mock instance
func NewMocktopicHandle(ctrl *gomock.Controller) *MocktopicHandle {
	mock := &MocktopicHandle{ctrl: ctrl}
	mock.recorder = &MocktopicHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocktopicHandle) EXPECT() *MocktopicHandleMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MocktopicHandle) Publish(ctx context.Context, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MocktopicHandleMockRecorder) Publish(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MocktopicHandle)(nil).Publish), ctx, data)
}

// Subscribe mocks base method
func (m *MocktopicHandle) Subscribe() (subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe")
	ret0, _ := ret[0].(subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MocktopicHandleMockRecorder) Subscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MocktopicHandle)(nil).Subscribe))
}

// MocktopicJoiner is a mock of topicJoiner interface
type MocktopicJoiner struct {
	ctrl     *gomock.Controller
	recorder *MocktopicJoinerMockRecorder
}

// MocktopicJoinerMockRecorder is the mock recorder for MocktopicJoiner
type MocktopicJoinerMockRecorder struct {
	mock *MocktopicJoiner
}

// NewMocktopicJoiner creates a new mock instance
func NewMocktopicJoiner(ctrl *gomock.Controller) *MocktopicJoiner {
	mock := &MocktopicJoiner{ctrl: ctrl}
	mock.recorder = &MocktopicJoinerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocktopicJoiner) EXPECT() *MocktopicJoinerMockRecorder {
	return m.recorder
}

// JoinTopic mocks base method
func (m *MocktopicJoiner) JoinTopic(topic string) (topicHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinTopic", topic)
	ret0, _ := ret[0].(topicHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JoinTopic indicates an expected call of JoinTopic
func (mr *MocktopicJoinerMockRecorder) JoinTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinTopic", reflect.TypeOf((*MocktopicJoiner)(nil).JoinTopic), topic)
}

// Mocksubscription is a mock of subscription interface
type Mocksubscription struct {
	ctrl     *gomock.Controller
	recorder *MocksubscriptionMockRecorder
}

// MocksubscriptionMockRecorder is the mock recorder for Mocksubscription
type MocksubscriptionMockRecorder struct {
	mock *Mocksubscription
}

// NewMocksubscription creates a new mock instance
func NewMocksubscription(ctrl *gomock.Controller) *Mocksubscription {
	mock := &Mocksubscription{ctrl: ctrl}
	mock.recorder = &MocksubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mocksubscription) EXPECT() *MocksubscriptionMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *Mocksubscription) Next(ctx context.Context) (*pubsub.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx)
	ret0, _ := ret[0].(*pubsub.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MocksubscriptionMockRecorder) Next(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*Mocksubscription)(nil).Next), ctx)
}

// Cancel mocks base method
func (m *Mocksubscription) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel
func (mr *MocksubscriptionMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*Mocksubscription)(nil).Cancel))
}
