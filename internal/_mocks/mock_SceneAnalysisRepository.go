// Code generated by mockery v2.51.1. DO NOT EDIT.

package _mocks

import (
	context "context"

	entity "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockSceneAnalysisRepository is an autogenerated mock type for the SceneAnalysisRepository type
type MockSceneAnalysisRepository struct {
	mock.Mock
}

type MockSceneAnalysisRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSceneAnalysisRepository) EXPECT() *MockSceneAnalysisRepository_Expecter {
	return &MockSceneAnalysisRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockSceneAnalysisRepository) Get(ctx context.Context, id string) (*entity.SceneAnalysis, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.SceneAnalysis
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.SceneAnalysis, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.SceneAnalysis); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.SceneAnalysis)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSceneAnalysisRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSceneAnalysisRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockSceneAnalysisRepository_Expecter) Get(ctx interface{}, id interface{}) *MockSceneAnalysisRepository_Get_Call {
	return &MockSceneAnalysisRepository_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockSceneAnalysisRepository_Get_Call) Run(run func(ctx context.Context, id string)) *MockSceneAnalysisRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSceneAnalysisRepository_Get_Call) Return(_a0 *entity.SceneAnalysis, _a1 error) *MockSceneAnalysisRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSceneAnalysisRepository_Get_Call) RunAndReturn(run func(context.Context, string) (*entity.SceneAnalysis, error)) *MockSceneAnalysisRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: ctx
func (_m *MockSceneAnalysisRepository) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSceneAnalysisRepository_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockSceneAnalysisRepository_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSceneAnalysisRepository_Expecter) Init(ctx interface{}) *MockSceneAnalysisRepository_Init_Call {
	return &MockSceneAnalysisRepository_Init_Call{Call: _e.mock.On("Init", ctx)}
}

func (_c *MockSceneAnalysisRepository_Init_Call) Run(run func(ctx context.Context)) *MockSceneAnalysisRepository_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSceneAnalysisRepository_Init_Call) Return(_a0 error) *MockSceneAnalysisRepository_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSceneAnalysisRepository_Init_Call) RunAndReturn(run func(context.Context) error) *MockSceneAnalysisRepository_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, analysis
func (_m *MockSceneAnalysisRepository) Save(ctx context.Context, analysis entity.SceneAnalysis) error {
	ret := _m.Called(ctx, analysis)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.SceneAnalysis) error); ok {
		r0 = rf(ctx, analysis)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSceneAnalysisRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockSceneAnalysisRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - analysis entity.SceneAnalysis
func (_e *MockSceneAnalysisRepository_Expecter) Save(ctx interface{}, analysis interface{}) *MockSceneAnalysisRepository_Save_Call {
	return &MockSceneAnalysisRepository_Save_Call{Call: _e.mock.On("Save", ctx, analysis)}
}

func (_c *MockSceneAnalysisRepository_Save_Call) Run(run func(ctx context.Context, analysis entity.SceneAnalysis)) *MockSceneAnalysisRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entity.SceneAnalysis))
	})
	return _c
}

func (_c *MockSceneAnalysisRepository_Save_Call) Return(_a0 error) *MockSceneAnalysisRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSceneAnalysisRepository_Save_Call) RunAndReturn(run func(context.Context, entity.SceneAnalysis) error) *MockSceneAnalysisRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSceneAnalysisRepository creates a new instance of MockSceneAnalysisRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSceneAnalysisRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSceneAnalysisRepository {
	mock := &MockSceneAnalysisRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
