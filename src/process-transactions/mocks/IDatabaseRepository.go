// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	entities "jucabet/stori-challenge/process-transactions/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// IDatabaseRepository is an autogenerated mock type for the IDatabaseRepository type
type IDatabaseRepository struct {
	mock.Mock
}

// SaveFileInfo provides a mock function with given fields: _a0
func (_m *IDatabaseRepository) SaveFileInfo(_a0 *entities.FileCharge) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.FileCharge) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveTransaction provides a mock function with given fields: _a0
func (_m *IDatabaseRepository) SaveTransaction(_a0 *entities.Transaction) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Transaction) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIDatabaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDatabaseRepository creates a new instance of IDatabaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDatabaseRepository(t mockConstructorTestingTNewIDatabaseRepository) *IDatabaseRepository {
	mock := &IDatabaseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
