// Copyright 2018 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package indexer

import (
	"context"
	"math/big"

	"bytes"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/getamis/sirius/log"
	"github.com/maichain/eth-indexer/common"
	"github.com/maichain/eth-indexer/model"
	"github.com/maichain/eth-indexer/store"
)

// New news an indexer service
func New(client EthClient, storeManager store.Manager) *indexer {
	return &indexer{
		client:  client,
		manager: storeManager,
	}
}

type indexer struct {
	client  EthClient
	manager store.Manager
}

func (idx *indexer) SyncToTarget(ctx context.Context, targetBlock int64) error {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Get local state from db
	header, stateBlock, err := idx.getLocalState()
	if err != nil {
		return err
	}

	if targetBlock <= header.Number {
		log.Error("Local block number is ahead of target block", "from", header.Number, "target", targetBlock)
		return errors.New(fmt.Sprintf("targetBlock should be greater than %d", header.Number))
	}

	_, err = idx.sync(childCtx, header.Number, header.Hash, targetBlock, stateBlock.Number)
	if err != nil {
		log.Error("Failed to sync from ethereum", "from", header.Number, "target", targetBlock, "err", err)
		return err
	}
	return nil
}

func (idx *indexer) Listen(ctx context.Context, ch chan *types.Header, syncMissingBlocks bool) error {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var lastBlockHeader *types.Header
	var stateBlock *model.StateBlock
	if syncMissingBlocks {
		// Get local state from db
		header, localState, err := idx.getLocalState()
		if err != nil {
			return err
		}

		// Get latest blocks from ethereum
		latestBlock, err := idx.client.BlockByNumber(childCtx, nil)
		if err != nil {
			log.Error("Failed to get latest header from ethereum", "err", err)
			return err
		}
		lastBlockHeader = latestBlock.Header()
		stateBlock = localState

		// Sync missing blocks from ethereum
		stateBlock, err = idx.sync(childCtx, header.Number, header.Hash, lastBlockHeader.Number.Int64(), stateBlock.Number)
		if err != nil {
			log.Error("Failed to sync to latest blocks from ethereum", "from", header.Number, "err", err)
			return err
		}
	}

	// Listen new channel events
	_, err := idx.client.SubscribeNewHead(childCtx, ch)
	if err != nil {
		log.Error("Failed to subscribe event for new header from ethereum", "err", err)
		return err
	}

	for {
		select {
		case head := <-ch:
			log.Trace("Got new header", "number", head.Number, "hash", common.HashHex(head.Hash()))
			if lastBlockHeader == nil {
				from := head.Number.Int64() - 1
				stateBlock, err = idx.sync(childCtx, from, []byte{}, head.Number.Int64(), from)
			} else if head.Number.Cmp(lastBlockHeader.Number) > 0 {
				stateBlock, err = idx.sync(childCtx, lastBlockHeader.Number.Int64(), lastBlockHeader.Hash().Bytes(), head.Number.Int64(), stateBlock.Number)
			} else {
				// TODO: check if head's TD is higher than the one we have locally. if so, reorg!
				log.Debug("Discarding older header", "number", head.Number)
			}
			if err != nil {
				log.Error("Failed to sync to blocks from ethereum", "from", lastBlockHeader.Number, "fromHash", common.HashHex(lastBlockHeader.Hash()), "to", head.Number.Int64(), "fromState", stateBlock.Number, "err", err)
				return err
			}
			lastBlockHeader = head
		case <-childCtx.Done():
			return childCtx.Err()
		}
	}
}

func (idx *indexer) getLocalState() (header *model.Header, stateBlock *model.StateBlock, err error) {
	// Get latest header from db
	header, err = idx.manager.LatestHeader()
	if err != nil {
		if common.NotFoundError(err) {
			log.Info("The header db is empty")
			header = &model.Header{
				Number: -1,
			}
			err = nil
		} else {
			log.Error("Failed to get latest header from db", "err", err)
			return
		}
	}

	// Get latest state block from db
	stateBlock, err = idx.manager.LatestStateBlock()
	if err != nil {
		if common.NotFoundError(err) {
			log.Info("The state db is empty")
			stateBlock = &model.StateBlock{
				Number: 0,
			}
			err = nil
		} else {
			log.Error("Failed to get latest state block from db", "err", err)
			return
		}
	}
	return
}

// sync syncs the blocks and header into database
func (idx *indexer) sync(ctx context.Context, from int64, fromHash []byte, to int64, fromStateBlock int64) (*model.StateBlock, error) {
	// Update existing blocks from ethereum to db
	for i := from + 1; i <= to; i++ {
		block, err := idx.client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			log.Error("Failed to get block from ethereum", "number", i, "err", err)
			return nil, err
		}

		if len(fromHash) > 0 && !bytes.Equal(block.ParentHash().Bytes(), fromHash) {
			if err = idx.reorg(ctx, block); err != nil {
				log.Error("Failed to reorg", "number", i, "hash", common.HashHex(block.Hash()), "err", err)
				return nil, err
			}
		}

		fromStateBlock, err = idx.addBlockData(ctx, block, fromStateBlock)
		if err != nil {
			log.Error("Failed to insert block locally", "number", i, "err", err)
			return nil, err
		}
		fromHash = block.Hash().Bytes()
	}
	return &model.StateBlock{
		Number: fromStateBlock,
	}, nil
}

func (idx *indexer) addBlockData(ctx context.Context, block *types.Block, fromStateBlock int64) (int64, error) {
	blockNumber := block.Number().Int64()
	var receipts []*types.Receipt
	for _, tx := range block.Transactions() {
		r, err := idx.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Error("Failed to get receipt from ethereum", "number", blockNumber, "tx", tx.Hash(), "err", err)
			return fromStateBlock, err
		}
		receipts = append(receipts, r)
	}

	err := idx.manager.InsertBlock(block, receipts)
	if err != nil {
		log.Error("Failed to insert block", "number", blockNumber, "err", err)
		return fromStateBlock, err
	}
	log.Trace("Inserted block", "number", blockNumber, "hash", common.HashHex(block.Hash()), "txs", len(block.Transactions()))

	// Get modified accounts
	// Noted: we skip dump block or get modified state error because the state db may not exist
	var dump *state.Dump
	isGenesis := blockNumber == 0
	if isGenesis {
		dump, err = idx.client.DumpBlock(ctx, 0)
		if err != nil {
			log.Warn("Failed to get state from ethereum, ignore it", "number", blockNumber, "err", err)
			return fromStateBlock, nil
		}
	} else {
		// This API is only supported on our customized geth.
		dump, err = idx.client.ModifiedAccountStatesByNumber(ctx, uint64(fromStateBlock), block.Number().Uint64())
		if err != nil {
			log.Warn("Failed to get modified accounts from ethereum, ignore it", "from", fromStateBlock, "to", blockNumber, "err", err)
			return fromStateBlock, nil
		}
	}

	// Update state db
	err = idx.manager.UpdateState(block, dump)
	if err != nil {
		log.Error("Failed to update state to database", "number", blockNumber, "err", err)
		return fromStateBlock, err
	}
	log.Trace("Inserted state", "number", blockNumber, "hash", common.HashHex(block.Hash()), "accounts", len(dump.Accounts))

	return blockNumber, nil
}

func (idx *indexer) reorg(ctx context.Context, block *types.Block) error {
	log.Trace("Reorg: tracing starts", "from", block.Number(), "hash", common.HashHex(block.Hash()))
	var blocks []*types.Block
	for {
		thisBlock, err := idx.client.BlockByHash(ctx, block.ParentHash())
		if err != nil || thisBlock == nil {
			log.Error("Reorg: failed to get block from ethereum", "hash", block.ParentHash().Hex(), "err", err)
			return err
		}
		block = thisBlock
		blocks = append(blocks, block)

		dbHeader, err := idx.manager.GetHeaderByNumber(block.Number().Int64() - 1)
		if err != nil || dbHeader == nil {
			log.Error("Reorg: failed to get header from local db", "number", block.Number().Int64()-1, "err", err)
			return err
		}

		if bytes.Equal(dbHeader.Hash, block.ParentHash().Bytes()) {
			break
		}
	}
	log.Trace("Reorg: tracing stops", "at", block.Number(), "hash", block.Hash().Hex())
	idx.manager.DeleteDataFromBlock(block.Number().Int64())

	// Get local state from db
	_, stateBlock, err := idx.getLocalState()
	if err != nil {
		return err
	}
	fromStateBlock := stateBlock.Number
	for i := len(blocks) - 1; i >= 0; i-- {
		block = blocks[i]
		fromStateBlock, err = idx.addBlockData(ctx, block, fromStateBlock)
		if err != nil {
			log.Error("reorg: failed to insert block data", "number", i, "err", err)
			return err
		}
	}
	log.Trace("Reorg: done", "at", block.Number(), "hash", block.Hash().Hex())
	return nil
}