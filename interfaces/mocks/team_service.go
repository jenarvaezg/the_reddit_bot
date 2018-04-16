// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import bson "gopkg.in/mgo.v2/bson"
import interfaces "github.com/jenarvaezg/MagicHub/interfaces"
import mock "github.com/stretchr/testify/mock"
import models "github.com/jenarvaezg/MagicHub/models"

// TeamService is an autogenerated mock type for the TeamService type
type TeamService struct {
	mock.Mock
}

// CreateTeam provides a mock function with given fields: userID, name, image, description
func (_m *TeamService) CreateTeam(userID bson.ObjectId, name string, image string, description string) (*models.Team, error) {
	ret := _m.Called(userID, name, image, description)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(bson.ObjectId, string, string, string) *models.Team); ok {
		r0 = rf(userID, name, image, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId, string, string, string) error); ok {
		r1 = rf(userID, name, image, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *TeamService) FindByID(id string) (*models.Team, error) {
	ret := _m.Called(id)

	var r0 *models.Team
	if rf, ok := ret.Get(0).(func(string) *models.Team); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFiltered provides a mock function with given fields: limit, offset, search
func (_m *TeamService) FindFiltered(limit int, offset int, search string) ([]*models.Team, error) {
	ret := _m.Called(limit, offset, search)

	var r0 []*models.Team
	if rf, ok := ret.Get(0).(func(int, int, string) []*models.Team); ok {
		r0 = rf(limit, offset, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRouteNameFromName provides a mock function with given fields: _a0
func (_m *TeamService) GetRouteNameFromName(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTeamAdmins provides a mock function with given fields: userID, team
func (_m *TeamService) GetTeamAdmins(userID bson.ObjectId, team *models.Team) ([]*models.User, error) {
	ret := _m.Called(userID, team)

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func(bson.ObjectId, *models.Team) []*models.User); ok {
		r0 = rf(userID, team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId, *models.Team) error); ok {
		r1 = rf(userID, team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamMembers provides a mock function with given fields: userID, team
func (_m *TeamService) GetTeamMembers(userID bson.ObjectId, team *models.Team) ([]*models.User, error) {
	ret := _m.Called(userID, team)

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func(bson.ObjectId, *models.Team) []*models.User); ok {
		r0 = rf(userID, team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId, *models.Team) error); ok {
		r1 = rf(userID, team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamMembersCount provides a mock function with given fields: team
func (_m *TeamService) GetTeamMembersCount(team *models.Team) (int, error) {
	ret := _m.Called(team)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.Team) int); ok {
		r0 = rf(team)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Team) error); ok {
		r1 = rf(team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnAllServicesRegistered provides a mock function with given fields: r
func (_m *TeamService) OnAllServicesRegistered(r interfaces.Registry) {
	_m.Called(r)
}
