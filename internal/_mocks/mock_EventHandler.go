// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"

	events "github.com/aws/aws-lambda-go/events"
	mock "github.com/stretchr/testify/mock"
)

// MockEventHandler is an autogenerated mock type for the EventHandler type
type MockEventHandler struct {
	mock.Mock
}

type MockEventHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventHandler) EXPECT() *MockEventHandler_Expecter {
	return &MockEventHandler_Expecter{mock: &_m.Mock}
}

// HandleEvent provides a mock function with given fields: ctx, s3Event
func (_m *MockEventHandler) HandleEvent(ctx context.Context, s3Event events.S3Event) error {
	ret := _m.Called(ctx, s3Event)

	if len(ret) == 0 {
		panic("no return value specified for HandleEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, events.S3Event) error); ok {
		r0 = rf(ctx, s3Event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventHandler_HandleEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleEvent'
type MockEventHandler_HandleEvent_Call struct {
	*mock.Call
}

// HandleEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - s3Event events.S3Event
func (_e *MockEventHandler_Expecter) HandleEvent(ctx interface{}, s3Event interface{}) *MockEventHandler_HandleEvent_Call {
	return &MockEventHandler_HandleEvent_Call{Call: _e.mock.On("HandleEvent", ctx, s3Event)}
}

func (_c *MockEventHandler_HandleEvent_Call) Run(run func(ctx context.Context, s3Event events.S3Event)) *MockEventHandler_HandleEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(events.S3Event))
	})
	return _c
}

func (_c *MockEventHandler_HandleEvent_Call) Return(_a0 error) *MockEventHandler_HandleEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEventHandler_HandleEvent_Call) RunAndReturn(run func(context.Context, events.S3Event) error) *MockEventHandler_HandleEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEventHandler creates a new instance of MockEventHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEventHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEventHandler {
	mock := &MockEventHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
