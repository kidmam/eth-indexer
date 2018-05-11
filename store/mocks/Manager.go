// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/maichain/eth-indexer/model"
import state "github.com/ethereum/go-ethereum/core/state"

import types "github.com/ethereum/go-ethereum/core/types"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// DeleteDataFromBlock provides a mock function with given fields: blockNumber
func (_m *Manager) DeleteDataFromBlock(blockNumber int64) error {
	ret := _m.Called(blockNumber)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(blockNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHeaderByNumber provides a mock function with given fields: number
func (_m *Manager) GetHeaderByNumber(number int64) (*model.Header, error) {
	ret := _m.Called(number)

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func(int64) *model.Header); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBlock provides a mock function with given fields: block, receipts
func (_m *Manager) InsertBlock(block *types.Block, receipts []*types.Receipt) error {
	ret := _m.Called(block, receipts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Block, []*types.Receipt) error); ok {
		r0 = rf(block, receipts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LatestHeader provides a mock function with given fields:
func (_m *Manager) LatestHeader() (*model.Header, error) {
	ret := _m.Called()

	var r0 *model.Header
	if rf, ok := ret.Get(0).(func() *model.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestStateBlock provides a mock function with given fields:
func (_m *Manager) LatestStateBlock() (*model.StateBlock, error) {
	ret := _m.Called()

	var r0 *model.StateBlock
	if rf, ok := ret.Get(0).(func() *model.StateBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.StateBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateState provides a mock function with given fields: block, accounts
func (_m *Manager) UpdateState(block *types.Block, accounts map[string]state.DumpDirtyAccount) error {
	ret := _m.Called(block, accounts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Block, map[string]state.DumpDirtyAccount) error); ok {
		r0 = rf(block, accounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
