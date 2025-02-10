// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"
	io "io"

	entity "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockScriptParser is an autogenerated mock type for the ScriptParser type
type MockScriptParser struct {
	mock.Mock
}

type MockScriptParser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScriptParser) EXPECT() *MockScriptParser_Expecter {
	return &MockScriptParser_Expecter{mock: &_m.Mock}
}

// ParseScript provides a mock function with given fields: ctx, reader
func (_m *MockScriptParser) ParseScript(ctx context.Context, reader io.Reader) (*entity.Script, error) {
	ret := _m.Called(ctx, reader)

	if len(ret) == 0 {
		panic("no return value specified for ParseScript")
	}

	var r0 *entity.Script
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) (*entity.Script, error)); ok {
		return rf(ctx, reader)
	}
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) *entity.Script); ok {
		r0 = rf(ctx, reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Script)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, io.Reader) error); ok {
		r1 = rf(ctx, reader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockScriptParser_ParseScript_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseScript'
type MockScriptParser_ParseScript_Call struct {
	*mock.Call
}

// ParseScript is a helper method to define mock.On call
//   - ctx context.Context
//   - reader io.Reader
func (_e *MockScriptParser_Expecter) ParseScript(ctx interface{}, reader interface{}) *MockScriptParser_ParseScript_Call {
	return &MockScriptParser_ParseScript_Call{Call: _e.mock.On("ParseScript", ctx, reader)}
}

func (_c *MockScriptParser_ParseScript_Call) Run(run func(ctx context.Context, reader io.Reader)) *MockScriptParser_ParseScript_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(io.Reader))
	})
	return _c
}

func (_c *MockScriptParser_ParseScript_Call) Return(_a0 *entity.Script, _a1 error) *MockScriptParser_ParseScript_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockScriptParser_ParseScript_Call) RunAndReturn(run func(context.Context, io.Reader) (*entity.Script, error)) *MockScriptParser_ParseScript_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScriptParser creates a new instance of MockScriptParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScriptParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScriptParser {
	mock := &MockScriptParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
