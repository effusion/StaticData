// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "StaticData/domain"
import mock "github.com/stretchr/testify/mock"

// CsvMappingRepository is an autogenerated mock type for the CsvMappingRepository type
type CsvMappingRepository struct {
	mock.Mock
}

// AllByKind provides a mock function with given fields: kind
func (_m *CsvMappingRepository) AllByKind(kind string) []*domain.CsvMapping {
	ret := _m.Called(kind)

	var r0 []*domain.CsvMapping
	if rf, ok := ret.Get(0).(func(string) []*domain.CsvMapping); ok {
		r0 = rf(kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.CsvMapping)
		}
	}

	return r0
}

// AllByKindAndParticipant provides a mock function with given fields: kind, participant
func (_m *CsvMappingRepository) AllByKindAndParticipant(kind string, participant int) []*domain.CsvMapping {
	ret := _m.Called(kind, participant)

	var r0 []*domain.CsvMapping
	if rf, ok := ret.Get(0).(func(string, int) []*domain.CsvMapping); ok {
		r0 = rf(kind, participant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.CsvMapping)
		}
	}

	return r0
}
