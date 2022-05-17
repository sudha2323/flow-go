// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	context "context"

	access "github.com/onflow/flow/protobuf/go/flow/access"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// AccessAPIServer is an autogenerated mock type for the AccessAPIServer type
type AccessAPIServer struct {
	mock.Mock
}

// ExecuteScriptAtBlockHeight provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) ExecuteScriptAtBlockHeight(_a0 context.Context, _a1 *access.ExecuteScriptAtBlockHeightRequest) (*access.ExecuteScriptResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtBlockHeightRequest) *access.ExecuteScriptResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtBlockHeightRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtBlockID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) ExecuteScriptAtBlockID(_a0 context.Context, _a1 *access.ExecuteScriptAtBlockIDRequest) (*access.ExecuteScriptResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtBlockIDRequest) *access.ExecuteScriptResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtBlockIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScriptAtLatestBlock provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) ExecuteScriptAtLatestBlock(_a0 context.Context, _a1 *access.ExecuteScriptAtLatestBlockRequest) (*access.ExecuteScriptResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.ExecuteScriptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.ExecuteScriptAtLatestBlockRequest) *access.ExecuteScriptResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecuteScriptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.ExecuteScriptAtLatestBlockRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetAccount(_a0 context.Context, _a1 *access.GetAccountRequest) (*access.GetAccountResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.GetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountRequest) *access.GetAccountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.GetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtBlockHeight provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetAccountAtBlockHeight(_a0 context.Context, _a1 *access.GetAccountAtBlockHeightRequest) (*access.AccountResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.AccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountAtBlockHeightRequest) *access.AccountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.AccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountAtBlockHeightRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAtLatestBlock provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetAccountAtLatestBlock(_a0 context.Context, _a1 *access.GetAccountAtLatestBlockRequest) (*access.AccountResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.AccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetAccountAtLatestBlockRequest) *access.AccountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.AccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetAccountAtLatestBlockRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHeight provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetBlockByHeight(_a0 context.Context, _a1 *access.GetBlockByHeightRequest) (*access.BlockResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockByHeightRequest) *access.BlockResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockByHeightRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetBlockByID(_a0 context.Context, _a1 *access.GetBlockByIDRequest) (*access.BlockResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockByIDRequest) *access.BlockResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByHeight provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetBlockHeaderByHeight(_a0 context.Context, _a1 *access.GetBlockHeaderByHeightRequest) (*access.BlockHeaderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockHeaderByHeightRequest) *access.BlockHeaderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockHeaderByHeightRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeaderByID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetBlockHeaderByID(_a0 context.Context, _a1 *access.GetBlockHeaderByIDRequest) (*access.BlockHeaderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetBlockHeaderByIDRequest) *access.BlockHeaderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetBlockHeaderByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionByID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetCollectionByID(_a0 context.Context, _a1 *access.GetCollectionByIDRequest) (*access.CollectionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.CollectionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetCollectionByIDRequest) *access.CollectionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.CollectionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetCollectionByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForBlockIDs provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetEventsForBlockIDs(_a0 context.Context, _a1 *access.GetEventsForBlockIDsRequest) (*access.EventsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.EventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetEventsForBlockIDsRequest) *access.EventsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetEventsForBlockIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsForHeightRange provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetEventsForHeightRange(_a0 context.Context, _a1 *access.GetEventsForHeightRangeRequest) (*access.EventsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.EventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetEventsForHeightRangeRequest) *access.EventsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetEventsForHeightRangeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionResultForBlockID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetExecutionResultForBlockID(_a0 context.Context, _a1 *access.GetExecutionResultForBlockIDRequest) (*access.ExecutionResultForBlockIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.ExecutionResultForBlockIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetExecutionResultForBlockIDRequest) *access.ExecutionResultForBlockIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ExecutionResultForBlockIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetExecutionResultForBlockIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlock provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetLatestBlock(_a0 context.Context, _a1 *access.GetLatestBlockRequest) (*access.BlockResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestBlockRequest) *access.BlockResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestBlockRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockHeader provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetLatestBlockHeader(_a0 context.Context, _a1 *access.GetLatestBlockHeaderRequest) (*access.BlockHeaderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.BlockHeaderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestBlockHeaderRequest) *access.BlockHeaderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.BlockHeaderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestBlockHeaderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestProtocolStateSnapshot provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetLatestProtocolStateSnapshot(_a0 context.Context, _a1 *access.GetLatestProtocolStateSnapshotRequest) (*access.ProtocolStateSnapshotResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.ProtocolStateSnapshotResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetLatestProtocolStateSnapshotRequest) *access.ProtocolStateSnapshotResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ProtocolStateSnapshotResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetLatestProtocolStateSnapshotRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkParameters provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetNetworkParameters(_a0 context.Context, _a1 *access.GetNetworkParametersRequest) (*access.GetNetworkParametersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.GetNetworkParametersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetNetworkParametersRequest) *access.GetNetworkParametersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.GetNetworkParametersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetNetworkParametersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetTransaction(_a0 context.Context, _a1 *access.GetTransactionRequest) (*access.TransactionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.TransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionRequest) *access.TransactionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResult provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetTransactionResult(_a0 context.Context, _a1 *access.GetTransactionRequest) (*access.TransactionResultResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.TransactionResultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionRequest) *access.TransactionResultResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResultByIndex provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetTransactionResultByIndex(_a0 context.Context, _a1 *access.GetTransactionByIndexRequest) (*access.TransactionResultResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.TransactionResultResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionByIndexRequest) *access.TransactionResultResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResultResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionByIndexRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionResultsByBlockID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetTransactionResultsByBlockID(_a0 context.Context, _a1 *access.GetTransactionsByBlockIDRequest) (*access.TransactionResultsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.TransactionResultsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionsByBlockIDRequest) *access.TransactionResultsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionResultsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionsByBlockIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByBlockID provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) GetTransactionsByBlockID(_a0 context.Context, _a1 *access.GetTransactionsByBlockIDRequest) (*access.TransactionsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.TransactionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.GetTransactionsByBlockIDRequest) *access.TransactionsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.TransactionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.GetTransactionsByBlockIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) Ping(_a0 context.Context, _a1 *access.PingRequest) (*access.PingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.PingRequest) *access.PingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.PingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.PingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: _a0, _a1
func (_m *AccessAPIServer) SendTransaction(_a0 context.Context, _a1 *access.SendTransactionRequest) (*access.SendTransactionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *access.SendTransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *access.SendTransactionRequest) *access.SendTransactionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.SendTransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *access.SendTransactionRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccessAPIServer creates a new instance of AccessAPIServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccessAPIServer(t testing.TB) *AccessAPIServer {
	mock := &AccessAPIServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
