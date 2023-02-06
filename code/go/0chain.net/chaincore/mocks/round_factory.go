// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	round "0chain.net/chaincore/round"
	mock "github.com/stretchr/testify/mock"
)

// RoundFactory is an autogenerated mock type for the RoundFactory type
type RoundFactory struct {
	mock.Mock
}

// CreateRoundF provides a mock function with given fields: roundNum
func (_m *RoundFactory) CreateRoundF(roundNum int64) round.RoundI {
	ret := _m.Called(roundNum)

	var r0 round.RoundI
	if rf, ok := ret.Get(0).(func(int64) round.RoundI); ok {
		r0 = rf(roundNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(round.RoundI)
		}
	}

	return r0
}

type mockConstructorTestingTNewRoundFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoundFactory creates a new instance of RoundFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoundFactory(t mockConstructorTestingTNewRoundFactory) *RoundFactory {
	mock := &RoundFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
