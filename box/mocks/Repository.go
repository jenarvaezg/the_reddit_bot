// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import bson "gopkg.in/mgo.v2/bson"
import mock "github.com/stretchr/testify/mock"
import models "github.com/jenarvaezg/MagicHub/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// FindFiltered provides a mock function with given fields: limit, offset, teamID
func (_m *Repository) FindFiltered(limit int, offset int, teamID bson.ObjectId) ([]*models.Box, error) {
	ret := _m.Called(limit, offset, teamID)

	var r0 []*models.Box
	if rf, ok := ret.Get(0).(func(int, int, bson.ObjectId) []*models.Box); ok {
		r0 = rf(limit, offset, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Box)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, bson.ObjectId) error); ok {
		r1 = rf(limit, offset, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
