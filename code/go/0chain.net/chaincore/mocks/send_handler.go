// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"

	node "0chain.net/chaincore/node"
	mock "github.com/stretchr/testify/mock"
)

// SendHandler is an autogenerated mock type for the SendHandler type
type SendHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, n
func (_m *SendHandler) Execute(ctx context.Context, n *node.Node) bool {
	ret := _m.Called(ctx, n)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) bool); ok {
		r0 = rf(ctx, n)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewSendHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewSendHandler creates a new instance of SendHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSendHandler(t mockConstructorTestingTNewSendHandler) *SendHandler {
	mock := &SendHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
