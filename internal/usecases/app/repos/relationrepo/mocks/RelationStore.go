// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	relationentity "antia/internal/entities/relationentity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RelationStore is an autogenerated mock type for the RelationStore type
type RelationStore struct {
	mock.Mock
}

// AddRelation provides a mock function with given fields: ctx, relation
func (_m *RelationStore) AddRelation(ctx context.Context, relation *relationentity.Relation) error {
	ret := _m.Called(ctx, relation)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *relationentity.Relation) error); ok {
		r0 = rf(ctx, relation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRelationByID provides a mock function with given fields: ctx, id
func (_m *RelationStore) DeleteRelationByID(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRelationByUserID provides a mock function with given fields: ctx, userID
func (_m *RelationStore) GetRelationByUserID(ctx context.Context, userID uint64) ([]*relationentity.NamedRelation, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*relationentity.NamedRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]*relationentity.NamedRelation, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []*relationentity.NamedRelation); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*relationentity.NamedRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRelationStore creates a new instance of RelationStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRelationStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *RelationStore {
	mock := &RelationStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}