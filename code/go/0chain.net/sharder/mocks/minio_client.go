// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	minio "github.com/minio/minio-go"
	mock "github.com/stretchr/testify/mock"
)

// MinioClient is an autogenerated mock type for the MinioClient type
type MinioClient struct {
	mock.Mock
}

// BucketExists provides a mock function with given fields: bucketName
func (_m *MinioClient) BucketExists(bucketName string) (bool, error) {
	ret := _m.Called(bucketName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(bucketName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucketName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BucketName provides a mock function with given fields:
func (_m *MinioClient) BucketName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DeleteLocal provides a mock function with given fields:
func (_m *MinioClient) DeleteLocal() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FGetObject provides a mock function with given fields: bucketName, objectName, filePath, options
func (_m *MinioClient) FGetObject(bucketName string, objectName string, filePath string, options minio.GetObjectOptions) error {
	ret := _m.Called(bucketName, objectName, filePath, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, minio.GetObjectOptions) error); ok {
		r0 = rf(bucketName, objectName, filePath, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FPutObject provides a mock function with given fields: bucketName, hash, filePath, options
func (_m *MinioClient) FPutObject(bucketName string, hash string, filePath string, options minio.PutObjectOptions) (int64, error) {
	ret := _m.Called(bucketName, hash, filePath, options)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, string, minio.PutObjectOptions) int64); ok {
		r0 = rf(bucketName, hash, filePath, options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, minio.PutObjectOptions) error); ok {
		r1 = rf(bucketName, hash, filePath, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeBucket provides a mock function with given fields: bucketName, location
func (_m *MinioClient) MakeBucket(bucketName string, location string) error {
	ret := _m.Called(bucketName, location)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bucketName, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StatObject provides a mock function with given fields: bucketName, hash, opts
func (_m *MinioClient) StatObject(bucketName string, hash string, opts minio.StatObjectOptions) (minio.ObjectInfo, error) {
	ret := _m.Called(bucketName, hash, opts)

	var r0 minio.ObjectInfo
	if rf, ok := ret.Get(0).(func(string, string, minio.StatObjectOptions) minio.ObjectInfo); ok {
		r0 = rf(bucketName, hash, opts)
	} else {
		r0 = ret.Get(0).(minio.ObjectInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, minio.StatObjectOptions) error); ok {
		r1 = rf(bucketName, hash, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMinioClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMinioClient creates a new instance of MinioClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMinioClient(t mockConstructorTestingTNewMinioClient) *MinioClient {
	mock := &MinioClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
