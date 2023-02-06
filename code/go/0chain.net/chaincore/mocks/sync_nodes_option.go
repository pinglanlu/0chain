// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	chain "0chain.net/chaincore/chain"
	mock "github.com/stretchr/testify/mock"
)

// SyncNodesOption is an autogenerated mock type for the SyncNodesOption type
type SyncNodesOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *SyncNodesOption) Execute(_a0 *chain.SyncReplyC) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewSyncNodesOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewSyncNodesOption creates a new instance of SyncNodesOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSyncNodesOption(t mockConstructorTestingTNewSyncNodesOption) *SyncNodesOption {
	mock := &SyncNodesOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
