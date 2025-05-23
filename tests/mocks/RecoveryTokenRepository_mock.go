// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/model"
	"github.com/CPU-commits/Template_Go-EventDriven/src/auth/repository/recovery_tokens_repository"
	mock "github.com/stretchr/testify/mock"
)

// NewMockRecoveryTokenRepository creates a new instance of MockRecoveryTokenRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRecoveryTokenRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRecoveryTokenRepository {
	mock := &MockRecoveryTokenRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockRecoveryTokenRepository is an autogenerated mock type for the RecoveryTokenRepository type
type MockRecoveryTokenRepository struct {
	mock.Mock
}

type MockRecoveryTokenRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRecoveryTokenRepository) EXPECT() *MockRecoveryTokenRepository_Expecter {
	return &MockRecoveryTokenRepository_Expecter{mock: &_m.Mock}
}

// Exists provides a mock function for the type MockRecoveryTokenRepository
func (_mock *MockRecoveryTokenRepository) Exists(criteria *recovery_tokens_repository.RecoveryTokenCriteria) (bool, error) {
	ret := _mock.Called(criteria)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) (bool, error)); ok {
		return returnFunc(criteria)
	}
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) bool); ok {
		r0 = returnFunc(criteria)
	} else {
		r0 = ret.Get(0).(bool)
	}
	if returnFunc, ok := ret.Get(1).(func(*recovery_tokens_repository.RecoveryTokenCriteria) error); ok {
		r1 = returnFunc(criteria)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRecoveryTokenRepository_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockRecoveryTokenRepository_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - criteria
func (_e *MockRecoveryTokenRepository_Expecter) Exists(criteria interface{}) *MockRecoveryTokenRepository_Exists_Call {
	return &MockRecoveryTokenRepository_Exists_Call{Call: _e.mock.On("Exists", criteria)}
}

func (_c *MockRecoveryTokenRepository_Exists_Call) Run(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria)) *MockRecoveryTokenRepository_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*recovery_tokens_repository.RecoveryTokenCriteria))
	})
	return _c
}

func (_c *MockRecoveryTokenRepository_Exists_Call) Return(b bool, err error) *MockRecoveryTokenRepository_Exists_Call {
	_c.Call.Return(b, err)
	return _c
}

func (_c *MockRecoveryTokenRepository_Exists_Call) RunAndReturn(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria) (bool, error)) *MockRecoveryTokenRepository_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function for the type MockRecoveryTokenRepository
func (_mock *MockRecoveryTokenRepository) Find(criteria *recovery_tokens_repository.RecoveryTokenCriteria) ([]model.RecoveryToken, error) {
	ret := _mock.Called(criteria)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []model.RecoveryToken
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) ([]model.RecoveryToken, error)); ok {
		return returnFunc(criteria)
	}
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) []model.RecoveryToken); ok {
		r0 = returnFunc(criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.RecoveryToken)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(*recovery_tokens_repository.RecoveryTokenCriteria) error); ok {
		r1 = returnFunc(criteria)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRecoveryTokenRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockRecoveryTokenRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - criteria
func (_e *MockRecoveryTokenRepository_Expecter) Find(criteria interface{}) *MockRecoveryTokenRepository_Find_Call {
	return &MockRecoveryTokenRepository_Find_Call{Call: _e.mock.On("Find", criteria)}
}

func (_c *MockRecoveryTokenRepository_Find_Call) Run(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria)) *MockRecoveryTokenRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*recovery_tokens_repository.RecoveryTokenCriteria))
	})
	return _c
}

func (_c *MockRecoveryTokenRepository_Find_Call) Return(recoveryTokens []model.RecoveryToken, err error) *MockRecoveryTokenRepository_Find_Call {
	_c.Call.Return(recoveryTokens, err)
	return _c
}

func (_c *MockRecoveryTokenRepository_Find_Call) RunAndReturn(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria) ([]model.RecoveryToken, error)) *MockRecoveryTokenRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function for the type MockRecoveryTokenRepository
func (_mock *MockRecoveryTokenRepository) FindOne(criteria *recovery_tokens_repository.RecoveryTokenCriteria) (*model.RecoveryToken, error) {
	ret := _mock.Called(criteria)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 *model.RecoveryToken
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) (*model.RecoveryToken, error)); ok {
		return returnFunc(criteria)
	}
	if returnFunc, ok := ret.Get(0).(func(*recovery_tokens_repository.RecoveryTokenCriteria) *model.RecoveryToken); ok {
		r0 = returnFunc(criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RecoveryToken)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(*recovery_tokens_repository.RecoveryTokenCriteria) error); ok {
		r1 = returnFunc(criteria)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRecoveryTokenRepository_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockRecoveryTokenRepository_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - criteria
func (_e *MockRecoveryTokenRepository_Expecter) FindOne(criteria interface{}) *MockRecoveryTokenRepository_FindOne_Call {
	return &MockRecoveryTokenRepository_FindOne_Call{Call: _e.mock.On("FindOne", criteria)}
}

func (_c *MockRecoveryTokenRepository_FindOne_Call) Run(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria)) *MockRecoveryTokenRepository_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*recovery_tokens_repository.RecoveryTokenCriteria))
	})
	return _c
}

func (_c *MockRecoveryTokenRepository_FindOne_Call) Return(recoveryToken *model.RecoveryToken, err error) *MockRecoveryTokenRepository_FindOne_Call {
	_c.Call.Return(recoveryToken, err)
	return _c
}

func (_c *MockRecoveryTokenRepository_FindOne_Call) RunAndReturn(run func(criteria *recovery_tokens_repository.RecoveryTokenCriteria) (*model.RecoveryToken, error)) *MockRecoveryTokenRepository_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// InsertOne provides a mock function for the type MockRecoveryTokenRepository
func (_mock *MockRecoveryTokenRepository) InsertOne(recoveryToken model.RecoveryToken) (*model.RecoveryToken, error) {
	ret := _mock.Called(recoveryToken)

	if len(ret) == 0 {
		panic("no return value specified for InsertOne")
	}

	var r0 *model.RecoveryToken
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(model.RecoveryToken) (*model.RecoveryToken, error)); ok {
		return returnFunc(recoveryToken)
	}
	if returnFunc, ok := ret.Get(0).(func(model.RecoveryToken) *model.RecoveryToken); ok {
		r0 = returnFunc(recoveryToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RecoveryToken)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(model.RecoveryToken) error); ok {
		r1 = returnFunc(recoveryToken)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRecoveryTokenRepository_InsertOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertOne'
type MockRecoveryTokenRepository_InsertOne_Call struct {
	*mock.Call
}

// InsertOne is a helper method to define mock.On call
//   - recoveryToken
func (_e *MockRecoveryTokenRepository_Expecter) InsertOne(recoveryToken interface{}) *MockRecoveryTokenRepository_InsertOne_Call {
	return &MockRecoveryTokenRepository_InsertOne_Call{Call: _e.mock.On("InsertOne", recoveryToken)}
}

func (_c *MockRecoveryTokenRepository_InsertOne_Call) Run(run func(recoveryToken model.RecoveryToken)) *MockRecoveryTokenRepository_InsertOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.RecoveryToken))
	})
	return _c
}

func (_c *MockRecoveryTokenRepository_InsertOne_Call) Return(recoveryToken1 *model.RecoveryToken, err error) *MockRecoveryTokenRepository_InsertOne_Call {
	_c.Call.Return(recoveryToken1, err)
	return _c
}

func (_c *MockRecoveryTokenRepository_InsertOne_Call) RunAndReturn(run func(recoveryToken model.RecoveryToken) (*model.RecoveryToken, error)) *MockRecoveryTokenRepository_InsertOne_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOne provides a mock function for the type MockRecoveryTokenRepository
func (_mock *MockRecoveryTokenRepository) UpdateOne(id int64, dataUpdate recovery_tokens_repository.RecoveryTokenUpdate) error {
	ret := _mock.Called(id, dataUpdate)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOne")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(int64, recovery_tokens_repository.RecoveryTokenUpdate) error); ok {
		r0 = returnFunc(id, dataUpdate)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockRecoveryTokenRepository_UpdateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOne'
type MockRecoveryTokenRepository_UpdateOne_Call struct {
	*mock.Call
}

// UpdateOne is a helper method to define mock.On call
//   - id
//   - dataUpdate
func (_e *MockRecoveryTokenRepository_Expecter) UpdateOne(id interface{}, dataUpdate interface{}) *MockRecoveryTokenRepository_UpdateOne_Call {
	return &MockRecoveryTokenRepository_UpdateOne_Call{Call: _e.mock.On("UpdateOne", id, dataUpdate)}
}

func (_c *MockRecoveryTokenRepository_UpdateOne_Call) Run(run func(id int64, dataUpdate recovery_tokens_repository.RecoveryTokenUpdate)) *MockRecoveryTokenRepository_UpdateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(recovery_tokens_repository.RecoveryTokenUpdate))
	})
	return _c
}

func (_c *MockRecoveryTokenRepository_UpdateOne_Call) Return(err error) *MockRecoveryTokenRepository_UpdateOne_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockRecoveryTokenRepository_UpdateOne_Call) RunAndReturn(run func(id int64, dataUpdate recovery_tokens_repository.RecoveryTokenUpdate) error) *MockRecoveryTokenRepository_UpdateOne_Call {
	_c.Call.Return(run)
	return _c
}
