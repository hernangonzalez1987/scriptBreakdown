// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockStorage is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

type MockStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorage) EXPECT() *MockStorage_Expecter {
	return &MockStorage_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, fileName, writer
func (_m *MockStorage) Get(ctx context.Context, fileName string, writer io.Writer) error {
	ret := _m.Called(ctx, fileName, writer)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Writer) error); ok {
		r0 = rf(ctx, fileName, writer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockStorage_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - fileName string
//   - writer io.Writer
func (_e *MockStorage_Expecter) Get(ctx interface{}, fileName interface{}, writer interface{}) *MockStorage_Get_Call {
	return &MockStorage_Get_Call{Call: _e.mock.On("Get", ctx, fileName, writer)}
}

func (_c *MockStorage_Get_Call) Run(run func(ctx context.Context, fileName string, writer io.Writer)) *MockStorage_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(io.Writer))
	})
	return _c
}

func (_c *MockStorage_Get_Call) Return(_a0 error) *MockStorage_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_Get_Call) RunAndReturn(run func(context.Context, string, io.Writer) error) *MockStorage_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, fileName, reader
func (_m *MockStorage) Put(ctx context.Context, fileName string, reader io.Reader) error {
	ret := _m.Called(ctx, fileName, reader)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader) error); ok {
		r0 = rf(ctx, fileName, reader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockStorage_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - fileName string
//   - reader io.Reader
func (_e *MockStorage_Expecter) Put(ctx interface{}, fileName interface{}, reader interface{}) *MockStorage_Put_Call {
	return &MockStorage_Put_Call{Call: _e.mock.On("Put", ctx, fileName, reader)}
}

func (_c *MockStorage_Put_Call) Run(run func(ctx context.Context, fileName string, reader io.Reader)) *MockStorage_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader))
	})
	return _c
}

func (_c *MockStorage_Put_Call) Return(_a0 error) *MockStorage_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_Put_Call) RunAndReturn(run func(context.Context, string, io.Reader) error) *MockStorage_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
