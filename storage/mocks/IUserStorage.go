// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/mado/model"
)

// IUserStorage is an autogenerated mock type for the IUserStorage type
type IUserStorage struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *IUserStorage) CreateUser(ctx context.Context, user model.User) (uint64, error) {
	ret := _m.Called(ctx, user)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) (uint64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User) uint64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, userID
func (_m *IUserStorage) DeleteUser(ctx context.Context, userID uint64) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, userID
func (_m *IUserStorage) GetUser(ctx context.Context, userID uint64) (model.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (model.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) model.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *IUserStorage) UpdateUser(ctx context.Context, user model.User) (uint64, error) {
	ret := _m.Called(ctx, user)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) (uint64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User) uint64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserStorage creates a new instance of IUserStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserStorage {
	mock := &IUserStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
