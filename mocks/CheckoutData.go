// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	checkout "warehouse/features/checkout"

	coreapi "github.com/midtrans/midtrans-go/coreapi"

	mock "github.com/stretchr/testify/mock"
)

// CheckoutData is an autogenerated mock type for the DataInterface type
type CheckoutData struct {
	mock.Mock
}

// AddCheckoutByFav provides a mock function with given fields: data
func (_m *CheckoutData) AddCheckoutByFav(data checkout.Core) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(checkout.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(checkout.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataPayment provides a mock function with given fields: reqPay
func (_m *CheckoutData) CreateDataPayment(reqPay coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	ret := _m.Called(reqPay)

	var r0 *coreapi.ChargeResponse
	if rf, ok := ret.Get(0).(func(coreapi.ChargeReq) *coreapi.ChargeResponse); ok {
		r0 = rf(reqPay)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coreapi.ChargeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coreapi.ChargeReq) error); ok {
		r1 = rf(reqPay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PaymentDataWebHook provides a mock function with given fields: data
func (_m *CheckoutData) PaymentDataWebHook(data checkout.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(checkout.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectPayment provides a mock function with given fields: orderID
func (_m *CheckoutData) SelectPayment(orderID string) (checkout.Core, error) {
	ret := _m.Called(orderID)

	var r0 checkout.Core
	if rf, ok := ret.Get(0).(func(string) checkout.Core); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Get(0).(checkout.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCheckoutData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCheckoutData creates a new instance of CheckoutData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCheckoutData(t mockConstructorTestingTNewCheckoutData) *CheckoutData {
	mock := &CheckoutData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
