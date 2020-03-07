// Code generated by mockery v1.0.0. DO NOT EDIT.

package parser

import mock "github.com/stretchr/testify/mock"

// MockParser is an autogenerated mock type for the Parser type
type MockParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields:
func (_m *MockParser) Parse() (*Service, error) {
	ret := _m.Called()

	var r0 *Service
	if rf, ok := ret.Get(0).(func() *Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
