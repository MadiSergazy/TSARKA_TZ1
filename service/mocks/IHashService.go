// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/mado/model"
)

// IHashService is an autogenerated mock type for the IHashService type
type IHashService struct {
	mock.Mock
}

// CalculateHash provides a mock function with given fields: certHash
func (_m *IHashService) CalculateHash(certHash model.CertHash) (uint64, error) {
	ret := _m.Called(certHash)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(model.CertHash) (uint64, error)); ok {
		return rf(certHash)
	}
	if rf, ok := ret.Get(0).(func(model.CertHash) uint64); ok {
		r0 = rf(certHash)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(model.CertHash) error); ok {
		r1 = rf(certHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCalculatedHash provides a mock function with given fields: hashID
func (_m *IHashService) GetCalculatedHash(hashID uint64) (model.CertHash, error) {
	ret := _m.Called(hashID)

	var r0 model.CertHash
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (model.CertHash, error)); ok {
		return rf(hashID)
	}
	if rf, ok := ret.Get(0).(func(uint64) model.CertHash); ok {
		r0 = rf(hashID)
	} else {
		r0 = ret.Get(0).(model.CertHash)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(hashID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIHashService creates a new instance of IHashService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHashService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHashService {
	mock := &IHashService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}