// Code generated by MockGen. DO NOT EDIT.
// Source: ../clock.go

// Package mockclock is a generated GoMock package.
package mockclock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	clock "github.com/transcelestial/clock"
)

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// NewTicker mocks base method.
func (m *MockClock) NewTicker(d time.Duration) clock.Ticker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTicker", d)
	ret0, _ := ret[0].(clock.Ticker)
	return ret0
}

// NewTicker indicates an expected call of NewTicker.
func (mr *MockClockMockRecorder) NewTicker(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTicker", reflect.TypeOf((*MockClock)(nil).NewTicker), d)
}

// NewTimer mocks base method.
func (m *MockClock) NewTimer(d time.Duration) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", d)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer.
func (mr *MockClockMockRecorder) NewTimer(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockClock)(nil).NewTimer), d)
}

// Now mocks base method.
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// Sleep mocks base method.
func (m *MockClock) Sleep(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep.
func (mr *MockClockMockRecorder) Sleep(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockClock)(nil).Sleep), d)
}