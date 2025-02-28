// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ITopicRepository is an autogenerated mock type for the ITopicRepository type
type ITopicRepository struct {
	mock.Mock
}

// SendMessageToReport provides a mock function with given fields: fileChargeID
func (_m *ITopicRepository) SendMessageToReport(fileChargeID string) error {
	ret := _m.Called(fileChargeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileChargeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewITopicRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewITopicRepository creates a new instance of ITopicRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITopicRepository(t mockConstructorTestingTNewITopicRepository) *ITopicRepository {
	mock := &ITopicRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
