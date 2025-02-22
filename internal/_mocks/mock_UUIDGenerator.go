// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockUUIDGenerator is an autogenerated mock type for the UUIDGenerator type
type MockUUIDGenerator struct {
	mock.Mock
}

type MockUUIDGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUUIDGenerator) EXPECT() *MockUUIDGenerator_Expecter {
	return &MockUUIDGenerator_Expecter{mock: &_m.Mock}
}

// New provides a mock function with no fields
func (_m *MockUUIDGenerator) New() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// MockUUIDGenerator_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type MockUUIDGenerator_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
func (_e *MockUUIDGenerator_Expecter) New() *MockUUIDGenerator_New_Call {
	return &MockUUIDGenerator_New_Call{Call: _e.mock.On("New")}
}

func (_c *MockUUIDGenerator_New_Call) Run(run func()) *MockUUIDGenerator_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUUIDGenerator_New_Call) Return(_a0 uuid.UUID) *MockUUIDGenerator_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUUIDGenerator_New_Call) RunAndReturn(run func() uuid.UUID) *MockUUIDGenerator_New_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUUIDGenerator creates a new instance of MockUUIDGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUUIDGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUUIDGenerator {
	mock := &MockUUIDGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
