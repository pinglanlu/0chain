// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package tokenpool

import mock "github.com/stretchr/testify/mock"

// MockTokenLockInterface is an autogenerated mock type for the TokenLockInterface type
type MockTokenLockInterface struct {
	mock.Mock
}

// IsLocked provides a mock function with given fields: entity
func (_m *MockTokenLockInterface) IsLocked(entity interface{}) bool {
	ret := _m.Called(entity)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LockStats provides a mock function with given fields: entity
func (_m *MockTokenLockInterface) LockStats(entity interface{}) []byte {
	ret := _m.Called(entity)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}) []byte); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MarshalMsg provides a mock function with given fields: _a0
func (_m *MockTokenLockInterface) MarshalMsg(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Msgsize provides a mock function with given fields:
func (_m *MockTokenLockInterface) Msgsize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UnmarshalMsg provides a mock function with given fields: _a0
func (_m *MockTokenLockInterface) UnmarshalMsg(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockTokenLockInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTokenLockInterface creates a new instance of MockTokenLockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTokenLockInterface(t mockConstructorTestingTNewMockTokenLockInterface) *MockTokenLockInterface {
	mock := &MockTokenLockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
