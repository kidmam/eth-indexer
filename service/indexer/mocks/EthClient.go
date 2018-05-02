// Code generated by mockery v1.0.0
package mocks

import big "math/big"
import common "github.com/ethereum/go-ethereum/common"
import context "context"
import ethereum "github.com/ethereum/go-ethereum"

import mock "github.com/stretchr/testify/mock"
import state "github.com/ethereum/go-ethereum/core/state"
import types "github.com/ethereum/go-ethereum/core/types"

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *EthClient) Close() {
	_m.Called()
}

// DumpBlock provides a mock function with given fields: ctx, blockNr
func (_m *EthClient) DumpBlock(ctx context.Context, blockNr int64) (*state.Dump, error) {
	ret := _m.Called(ctx, blockNr)

	var r0 *state.Dump
	if rf, ok := ret.Get(0).(func(context.Context, int64) *state.Dump); ok {
		r0 = rf(ctx, blockNr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Dump)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, blockNr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifiedAccountStatesByNumber provides a mock function with given fields: ctx, startNum, endNum
func (_m *EthClient) ModifiedAccountStatesByNumber(ctx context.Context, startNum uint64, endNum uint64) (*state.Dump, error) {
	ret := _m.Called(ctx, startNum, endNum)

	var r0 *state.Dump
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) *state.Dump); ok {
		r0 = rf(ctx, startNum, endNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Dump)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64) error); ok {
		r1 = rf(ctx, startNum, endNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *EthClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	var r0 ethereum.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *EthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}