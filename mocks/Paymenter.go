// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gin "github.com/gin-gonic/gin"
import mock "github.com/stretchr/testify/mock"

// Paymenter is an autogenerated mock type for the Paymenter type
type Paymenter struct {
	mock.Mock
}

// AddPayment provides a mock function with given fields: ctx
func (_m *Paymenter) AddPayment(ctx *gin.Context) {
	_m.Called(ctx)
}
