// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domains "github.com/lucaspizzo/api-payment/domains"
import forms "github.com/lucaspizzo/api-payment/forms"
import mock "github.com/stretchr/testify/mock"

// AccountContract is an autogenerated mock type for the AccountContract type
type AccountContract struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *AccountContract) Get(id uint64) (*domains.Account, error) {
	ret := _m.Called(id)

	var r0 *domains.Account
	if rf, ok := ret.Get(0).(func(uint64) *domains.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *AccountContract) List() (*[]*domains.Account, error) {
	ret := _m.Called()

	var r0 *[]*domains.Account
	if rf, ok := ret.Get(0).(func() *[]*domains.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]*domains.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: availableCreditLimit, availableWithdrawalLimit, account
func (_m *AccountContract) Update(availableCreditLimit float64, availableWithdrawalLimit float64, account *domains.Account) (*domains.Account, error) {
	ret := _m.Called(availableCreditLimit, availableWithdrawalLimit, account)

	var r0 *domains.Account
	if rf, ok := ret.Get(0).(func(float64, float64, *domains.Account) *domains.Account); ok {
		r0 = rf(availableCreditLimit, availableWithdrawalLimit, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float64, float64, *domains.Account) error); ok {
		r1 = rf(availableCreditLimit, availableWithdrawalLimit, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLimits provides a mock function with given fields: form
func (_m *AccountContract) UpdateLimits(form *forms.LimitForm) (*domains.Account, error) {
	ret := _m.Called(form)

	var r0 *domains.Account
	if rf, ok := ret.Get(0).(func(*forms.LimitForm) *domains.Account); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*forms.LimitForm) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}