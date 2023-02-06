// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	conductrpc "0chain.net/conductor/conductrpc"
	mock "github.com/stretchr/testify/mock"
)

// UpdateStateFunc is an autogenerated mock type for the UpdateStateFunc type
type UpdateStateFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: state
func (_m *UpdateStateFunc) Execute(state *conductrpc.State) {
	_m.Called(state)
}

type mockConstructorTestingTNewUpdateStateFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewUpdateStateFunc creates a new instance of UpdateStateFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUpdateStateFunc(t mockConstructorTestingTNewUpdateStateFunc) *UpdateStateFunc {
	mock := &UpdateStateFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
