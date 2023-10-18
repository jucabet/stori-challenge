// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
)

// ISQSInterface is an autogenerated mock type for the ISQSInterface type
type ISQSInterface struct {
	mock.Mock
}

// DeleteMessage provides a mock function with given fields: ctx, params, optFns
func (_m *ISQSInterface) DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sqs.DeleteMessageOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sqs.DeleteMessageInput, ...func(*sqs.Options)) *sqs.DeleteMessageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.DeleteMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqs.DeleteMessageInput, ...func(*sqs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQueueUrl provides a mock function with given fields: ctx, params, optFns
func (_m *ISQSInterface) GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sqs.GetQueueUrlOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sqs.GetQueueUrlInput, ...func(*sqs.Options)) *sqs.GetQueueUrlOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.GetQueueUrlOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqs.GetQueueUrlInput, ...func(*sqs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiveMessage provides a mock function with given fields: ctx, params, optFns
func (_m *ISQSInterface) ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sqs.ReceiveMessageOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sqs.ReceiveMessageInput, ...func(*sqs.Options)) *sqs.ReceiveMessageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ReceiveMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqs.ReceiveMessageInput, ...func(*sqs.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewISQSInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewISQSInterface creates a new instance of ISQSInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISQSInterface(t mockConstructorTestingTNewISQSInterface) *ISQSInterface {
	mock := &ISQSInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
