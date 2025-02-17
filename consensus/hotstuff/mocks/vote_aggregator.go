// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"
	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"

	testing "testing"
)

// VoteAggregator is an autogenerated mock type for the VoteAggregator type
type VoteAggregator struct {
	mock.Mock
}

// AddBlock provides a mock function with given fields: block
func (_m *VoteAggregator) AddBlock(block *model.Proposal) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Proposal) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddVote provides a mock function with given fields: vote
func (_m *VoteAggregator) AddVote(vote *model.Vote) {
	_m.Called(vote)
}

// Done provides a mock function with given fields:
func (_m *VoteAggregator) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// InvalidBlock provides a mock function with given fields: block
func (_m *VoteAggregator) InvalidBlock(block *model.Proposal) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Proposal) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PruneUpToView provides a mock function with given fields: view
func (_m *VoteAggregator) PruneUpToView(view uint64) {
	_m.Called(view)
}

// Ready provides a mock function with given fields:
func (_m *VoteAggregator) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *VoteAggregator) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

// NewVoteAggregator creates a new instance of VoteAggregator. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewVoteAggregator(t testing.TB) *VoteAggregator {
	mock := &VoteAggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
