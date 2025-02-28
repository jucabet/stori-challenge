// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	dtos "jucabet/stori-challenge/send-reports/internal/domain/dtos"

	mock "github.com/stretchr/testify/mock"
)

// IStategiesInterface is an autogenerated mock type for the IStategiesInterface type
type IStategiesInterface struct {
	mock.Mock
}

// BuildReport provides a mock function with given fields: data
func (_m *IStategiesInterface) BuildReport(data *dtos.SendReportDto) (map[string]interface{}, error) {
	ret := _m.Called(data)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(*dtos.SendReportDto) map[string]interface{}); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dtos.SendReportDto) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendReport provides a mock function with given fields: reportData
func (_m *IStategiesInterface) SendReport(reportData map[string]interface{}) error {
	ret := _m.Called(reportData)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) error); ok {
		r0 = rf(reportData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIStategiesInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewIStategiesInterface creates a new instance of IStategiesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIStategiesInterface(t mockConstructorTestingTNewIStategiesInterface) *IStategiesInterface {
	mock := &IStategiesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
