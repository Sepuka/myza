// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CacheResponse is an autogenerated mock type for the CacheResponse type
type CacheResponse struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *CacheResponse) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Val provides a mock function with given fields:
func (_m *CacheResponse) Val() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}