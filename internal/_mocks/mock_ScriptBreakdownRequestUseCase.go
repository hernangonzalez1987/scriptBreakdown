// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"

	entity "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockScriptBreakdownRequestUseCase is an autogenerated mock type for the ScriptBreakdownRequestUseCase type
type MockScriptBreakdownRequestUseCase struct {
	mock.Mock
}

type MockScriptBreakdownRequestUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScriptBreakdownRequestUseCase) EXPECT() *MockScriptBreakdownRequestUseCase_Expecter {
	return &MockScriptBreakdownRequestUseCase_Expecter{mock: &_m.Mock}
}

// GetResult provides a mock function with given fields: ctx, breakdownID
func (_m *MockScriptBreakdownRequestUseCase) GetResult(ctx context.Context, breakdownID string) (*entity.ScriptBreakdownResult, error) {
	ret := _m.Called(ctx, breakdownID)

	if len(ret) == 0 {
		panic("no return value specified for GetResult")
	}

	var r0 *entity.ScriptBreakdownResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.ScriptBreakdownResult, error)); ok {
		return rf(ctx, breakdownID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.ScriptBreakdownResult); ok {
		r0 = rf(ctx, breakdownID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ScriptBreakdownResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, breakdownID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockScriptBreakdownRequestUseCase_GetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResult'
type MockScriptBreakdownRequestUseCase_GetResult_Call struct {
	*mock.Call
}

// GetResult is a helper method to define mock.On call
//   - ctx context.Context
//   - breakdownID string
func (_e *MockScriptBreakdownRequestUseCase_Expecter) GetResult(ctx interface{}, breakdownID interface{}) *MockScriptBreakdownRequestUseCase_GetResult_Call {
	return &MockScriptBreakdownRequestUseCase_GetResult_Call{Call: _e.mock.On("GetResult", ctx, breakdownID)}
}

func (_c *MockScriptBreakdownRequestUseCase_GetResult_Call) Run(run func(ctx context.Context, breakdownID string)) *MockScriptBreakdownRequestUseCase_GetResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockScriptBreakdownRequestUseCase_GetResult_Call) Return(_a0 *entity.ScriptBreakdownResult, _a1 error) *MockScriptBreakdownRequestUseCase_GetResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockScriptBreakdownRequestUseCase_GetResult_Call) RunAndReturn(run func(context.Context, string) (*entity.ScriptBreakdownResult, error)) *MockScriptBreakdownRequestUseCase_GetResult_Call {
	_c.Call.Return(run)
	return _c
}

// RequestScriptBreakdown provides a mock function with given fields: ctx, req
func (_m *MockScriptBreakdownRequestUseCase) RequestScriptBreakdown(ctx context.Context, req entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for RequestScriptBreakdown")
	}

	var r0 *entity.ScriptBreakdownResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ScriptBreakdownRequest) *entity.ScriptBreakdownResult); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.ScriptBreakdownResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ScriptBreakdownRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestScriptBreakdown'
type MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call struct {
	*mock.Call
}

// RequestScriptBreakdown is a helper method to define mock.On call
//   - ctx context.Context
//   - req entity.ScriptBreakdownRequest
func (_e *MockScriptBreakdownRequestUseCase_Expecter) RequestScriptBreakdown(ctx interface{}, req interface{}) *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call {
	return &MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call{Call: _e.mock.On("RequestScriptBreakdown", ctx, req)}
}

func (_c *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call) Run(run func(ctx context.Context, req entity.ScriptBreakdownRequest)) *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.ScriptBreakdownRequest))
	})
	return _c
}

func (_c *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call) Return(_a0 *entity.ScriptBreakdownResult, _a1 error) *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call) RunAndReturn(run func(context.Context, entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error)) *MockScriptBreakdownRequestUseCase_RequestScriptBreakdown_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScriptBreakdownRequestUseCase creates a new instance of MockScriptBreakdownRequestUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScriptBreakdownRequestUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScriptBreakdownRequestUseCase {
	mock := &MockScriptBreakdownRequestUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
