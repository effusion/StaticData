// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "StaticData/domain"
import gorm "github.com/jinzhu/gorm"
import mock "github.com/stretchr/testify/mock"

// AuditRepository is an autogenerated mock type for the AuditRepository type
type AuditRepository struct {
	mock.Mock
}

// SaveAudit provides a mock function with given fields: transaction
func (_m *AuditRepository) SaveAudit(transaction domain.AuditTransaction) gorm.Errors {
	ret := _m.Called(transaction)

	var r0 gorm.Errors
	if rf, ok := ret.Get(0).(func(domain.AuditTransaction) gorm.Errors); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gorm.Errors)
		}
	}

	return r0
}
