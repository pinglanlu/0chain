// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	encryption "0chain.net/core/encryption"
	mock "github.com/stretchr/testify/mock"
)

// ReconstructSignatureScheme is an autogenerated mock type for the ReconstructSignatureScheme type
type ReconstructSignatureScheme struct {
	mock.Mock
}

// Add provides a mock function with given fields: tss, signature
func (_m *ReconstructSignatureScheme) Add(tss encryption.ThresholdSignatureScheme, signature string) error {
	ret := _m.Called(tss, signature)

	var r0 error
	if rf, ok := ret.Get(0).(func(encryption.ThresholdSignatureScheme, string) error); ok {
		r0 = rf(tss, signature)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reconstruct provides a mock function with given fields:
func (_m *ReconstructSignatureScheme) Reconstruct() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReconstructSignatureScheme interface {
	mock.TestingT
	Cleanup(func())
}

// NewReconstructSignatureScheme creates a new instance of ReconstructSignatureScheme. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReconstructSignatureScheme(t mockConstructorTestingTNewReconstructSignatureScheme) *ReconstructSignatureScheme {
	mock := &ReconstructSignatureScheme{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
