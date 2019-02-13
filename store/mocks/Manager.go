// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import client "github.com/getamis/eth-indexer/client"
import common "github.com/ethereum/go-ethereum/common"
import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"

import types "github.com/ethereum/go-ethereum/core/types"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// FindBlockByNumber provides a mock function with given fields: ctx, number
func (_m *Manager) FindBlockByNumber(ctx context.Context, number int64) (*model.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindERC20 provides a mock function with given fields: ctx, address
func (_m *Manager) FindERC20(ctx context.Context, address common.Address) (*model.ERC20, error) {
	ret := _m.Called(ctx, address)

	var r0 *model.ERC20
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) *model.ERC20); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ERC20)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLatestBlock provides a mock function with given fields: ctx
func (_m *Manager) FindLatestBlock(ctx context.Context) (*model.Header, error) {
	ret := _m.Called(ctx)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func(context.Context) *model.Header); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTd provides a mock function with given fields: ctx, hash
func (_m *Manager) FindTd(ctx context.Context, hash []byte) (*model.TotalDifficulty, error) {
	ret := _m.Called(ctx, hash)

	var r0 *model.TotalDifficulty
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *model.TotalDifficulty); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TotalDifficulty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx
func (_m *Manager) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertBlocks provides a mock function with given fields: ctx, balancer, blocks, receipts, events
func (_m *Manager) InsertBlocks(ctx context.Context, balancer client.Balancer, blocks []*types.Block, receipts [][]*types.Receipt, events [][]*types.TransferLog) error {
	ret := _m.Called(ctx, balancer, blocks, receipts, events)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, client.Balancer, []*types.Block, [][]*types.Receipt, [][]*types.TransferLog) error); ok {
		r0 = rf(ctx, balancer, blocks, receipts, events)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertERC20 provides a mock function with given fields: ctx, code
func (_m *Manager) InsertERC20(ctx context.Context, code *model.ERC20) error {
	ret := _m.Called(ctx, code)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ERC20) error); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertTd provides a mock function with given fields: ctx, data
func (_m *Manager) InsertTd(ctx context.Context, data *model.TotalDifficulty) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TotalDifficulty) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReorgBlocks provides a mock function with given fields: ctx, reorgEvent
func (_m *Manager) ReorgBlocks(ctx context.Context, reorgEvent *model.Reorg) error {
	ret := _m.Called(ctx, reorgEvent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reorg) error); ok {
		r0 = rf(ctx, reorgEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
