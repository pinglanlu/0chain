// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BatchI is an autogenerated mock type for the BatchI type
type BatchI struct {
	mock.Mock
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *BatchI) Query(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewBatchI interface {
	mock.TestingT
	Cleanup(func())
}

// NewBatchI creates a new instance of BatchI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBatchI(t mockConstructorTestingTNewBatchI) *BatchI {
	mock := &BatchI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
