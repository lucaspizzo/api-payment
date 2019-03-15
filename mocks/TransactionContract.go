// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domains "github.com/lucaspizzo/api-payment/domains"
import forms "github.com/lucaspizzo/api-payment/forms"
import mock "github.com/stretchr/testify/mock"

// TransactionContract is an autogenerated mock type for the TransactionContract type
type TransactionContract struct {
	mock.Mock
}

// List provides a mock function with given fields:
func (_m *TransactionContract) List() (*[]*domains.Transaction, error) {
	ret := _m.Called()

	var r0 *[]*domains.Transaction
	if rf, ok := ret.Get(0).(func() *[]*domains.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]*domains.Transaction)
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

// ListByAccountId provides a mock function with given fields: accountID
func (_m *TransactionContract) ListByAccountId(accountID uint64) (*[]*domains.Transaction, error) {
	ret := _m.Called(accountID)

	var r0 *[]*domains.Transaction
	if rf, ok := ret.Get(0).(func(uint64) *[]*domains.Transaction); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]*domains.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterTransaction provides a mock function with given fields: form
func (_m *TransactionContract) RegisterTransaction(form *forms.TransactionForm) (*domains.Transaction, error) {
	ret := _m.Called(form)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(*forms.TransactionForm) *domains.Transaction); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*forms.TransactionForm) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: transaction
func (_m *TransactionContract) Update(transaction *domains.Transaction) (*domains.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(*domains.Transaction) *domains.Transaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domains.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// createPayment provides a mock function with given fields: account, operationType
func (_m *TransactionContract) createPayment(account *domains.Account, operationType *domains.OperationType) *domains.Transaction {
	ret := _m.Called(account, operationType)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(*domains.Account, *domains.OperationType) *domains.Transaction); ok {
		r0 = rf(account, operationType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	return r0
}

// process provides a mock function with given fields: accountID, operationTypeId, amount
func (_m *TransactionContract) process(accountID uint64, operationTypeId uint64, amount float64) (*domains.Transaction, error) {
	ret := _m.Called(accountID, operationTypeId, amount)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(uint64, uint64, float64) *domains.Transaction); ok {
		r0 = rf(accountID, operationTypeId, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64, float64) error); ok {
		r1 = rf(accountID, operationTypeId, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// processPayment provides a mock function with given fields: account, operationType, amount
func (_m *TransactionContract) processPayment(account *domains.Account, operationType *domains.OperationType, amount float64) *domains.Transaction {
	ret := _m.Called(account, operationType, amount)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(*domains.Account, *domains.OperationType, float64) *domains.Transaction); ok {
		r0 = rf(account, operationType, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	return r0
}

// processPurchase provides a mock function with given fields: account, operationType, amount
func (_m *TransactionContract) processPurchase(account *domains.Account, operationType *domains.OperationType, amount float64) *domains.Transaction {
	ret := _m.Called(account, operationType, amount)

	var r0 *domains.Transaction
	if rf, ok := ret.Get(0).(func(*domains.Account, *domains.OperationType, float64) *domains.Transaction); ok {
		r0 = rf(account, operationType, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Transaction)
		}
	}

	return r0
}