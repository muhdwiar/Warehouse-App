// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	favorite "warehouse/features/favorite"

	mock "github.com/stretchr/testify/mock"
)

// FavoriteData is an autogenerated mock type for the DataInterface type
type FavoriteData struct {
	mock.Mock
}

// AddFavorite provides a mock function with given fields: data
func (_m *FavoriteData) AddFavorite(data favorite.Core) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(favorite.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(favorite.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteData provides a mock function with given fields: token, idfav
func (_m *FavoriteData) DeleteData(token int, idfav int) (int, error) {
	ret := _m.Called(token, idfav)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(token, idfav)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(token, idfav)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectFavorite provides a mock function with given fields: token
func (_m *FavoriteData) SelectFavorite(token int) ([]favorite.Core, error) {
	ret := _m.Called(token)

	var r0 []favorite.Core
	if rf, ok := ret.Get(0).(func(int) []favorite.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]favorite.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFavoriteData interface {
	mock.TestingT
	Cleanup(func())
}

// NewFavoriteData creates a new instance of FavoriteData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFavoriteData(t mockConstructorTestingTNewFavoriteData) *FavoriteData {
	mock := &FavoriteData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
