// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	context "context"

	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"

	testing "testing"
)

// MutableState is an autogenerated mock type for the MutableState type
type MutableState struct {
	mock.Mock
}

// AtBlockID provides a mock function with given fields: blockID
func (_m *MutableState) AtBlockID(blockID flow.Identifier) protocol.Snapshot {
	ret := _m.Called(blockID)

	var r0 protocol.Snapshot
	if rf, ok := ret.Get(0).(func(flow.Identifier) protocol.Snapshot); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Snapshot)
		}
	}

	return r0
}

// AtHeight provides a mock function with given fields: height
func (_m *MutableState) AtHeight(height uint64) protocol.Snapshot {
	ret := _m.Called(height)

	var r0 protocol.Snapshot
	if rf, ok := ret.Get(0).(func(uint64) protocol.Snapshot); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Snapshot)
		}
	}

	return r0
}

// Extend provides a mock function with given fields: ctx, candidate
func (_m *MutableState) Extend(ctx context.Context, candidate *flow.Block) error {
	ret := _m.Called(ctx, candidate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *flow.Block) error); ok {
		r0 = rf(ctx, candidate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Final provides a mock function with given fields:
func (_m *MutableState) Final() protocol.Snapshot {
	ret := _m.Called()

	var r0 protocol.Snapshot
	if rf, ok := ret.Get(0).(func() protocol.Snapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Snapshot)
		}
	}

	return r0
}

// Finalize provides a mock function with given fields: ctx, blockID
func (_m *MutableState) Finalize(ctx context.Context, blockID flow.Identifier) error {
	ret := _m.Called(ctx, blockID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier) error); ok {
		r0 = rf(ctx, blockID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkValid provides a mock function with given fields: blockID
func (_m *MutableState) MarkValid(blockID flow.Identifier) error {
	ret := _m.Called(blockID)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) error); ok {
		r0 = rf(blockID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Params provides a mock function with given fields:
func (_m *MutableState) Params() protocol.Params {
	ret := _m.Called()

	var r0 protocol.Params
	if rf, ok := ret.Get(0).(func() protocol.Params); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Params)
		}
	}

	return r0
}

// Sealed provides a mock function with given fields:
func (_m *MutableState) Sealed() protocol.Snapshot {
	ret := _m.Called()

	var r0 protocol.Snapshot
	if rf, ok := ret.Get(0).(func() protocol.Snapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Snapshot)
		}
	}

	return r0
}

// NewMutableState creates a new instance of MutableState. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMutableState(t testing.TB) *MutableState {
	mock := &MutableState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
