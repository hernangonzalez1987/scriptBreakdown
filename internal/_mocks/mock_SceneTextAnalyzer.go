// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockSceneTextAnalyzer is an autogenerated mock type for the SceneTextAnalyzer type
type MockSceneTextAnalyzer struct {
	mock.Mock
}

type MockSceneTextAnalyzer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSceneTextAnalyzer) EXPECT() *MockSceneTextAnalyzer_Expecter {
	return &MockSceneTextAnalyzer_Expecter{mock: &_m.Mock}
}

// AnalyzeSceneText provides a mock function with given fields: ctx, sceneText
func (_m *MockSceneTextAnalyzer) AnalyzeSceneText(ctx context.Context, sceneText string) (map[string][]string, error) {
	ret := _m.Called(ctx, sceneText)

	if len(ret) == 0 {
		panic("no return value specified for AnalyzeSceneText")
	}

	var r0 map[string][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string][]string, error)); ok {
		return rf(ctx, sceneText)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string][]string); ok {
		r0 = rf(ctx, sceneText)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sceneText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSceneTextAnalyzer_AnalyzeSceneText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AnalyzeSceneText'
type MockSceneTextAnalyzer_AnalyzeSceneText_Call struct {
	*mock.Call
}

// AnalyzeSceneText is a helper method to define mock.On call
//   - ctx context.Context
//   - sceneText string
func (_e *MockSceneTextAnalyzer_Expecter) AnalyzeSceneText(ctx interface{}, sceneText interface{}) *MockSceneTextAnalyzer_AnalyzeSceneText_Call {
	return &MockSceneTextAnalyzer_AnalyzeSceneText_Call{Call: _e.mock.On("AnalyzeSceneText", ctx, sceneText)}
}

func (_c *MockSceneTextAnalyzer_AnalyzeSceneText_Call) Run(run func(ctx context.Context, sceneText string)) *MockSceneTextAnalyzer_AnalyzeSceneText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSceneTextAnalyzer_AnalyzeSceneText_Call) Return(_a0 map[string][]string, _a1 error) *MockSceneTextAnalyzer_AnalyzeSceneText_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSceneTextAnalyzer_AnalyzeSceneText_Call) RunAndReturn(run func(context.Context, string) (map[string][]string, error)) *MockSceneTextAnalyzer_AnalyzeSceneText_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSceneTextAnalyzer creates a new instance of MockSceneTextAnalyzer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSceneTextAnalyzer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSceneTextAnalyzer {
	mock := &MockSceneTextAnalyzer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
