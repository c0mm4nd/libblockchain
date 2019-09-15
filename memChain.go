package libblockchain

import (
	"bytes"
	"errors"
	"math/big"
	"sync"

	"github.com/maoxs2/libblockchain/abstract"
)

type MemChain struct {
	latestBlock  *Block
	mu           *sync.RWMutex
	memHashMap   map[string][]*Block
	memHeightMap map[string]*Block
}

var (
	ErrWrongPrevHash      = errors.New("wrong PrevHash")
	ErrHasSameHeightBlock = errors.New("the storage already has a block with the same GetHeight")
	ErrHasSameHashBlock   = errors.New("the storage already has a block with the same Hash")
	ErrMissingLatestBlock = errors.New("the latest block point is missing")
)

var (
	big1 = big.NewInt(1)
	big0 = big.NewInt(0)
)

// Push is the only legal method to add element to the chain
func (c *MemChain) Push(block abstract.Block) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var strHash string
	var hexHeight string
	var err error
	if c.latestBlock == nil {
		if len(c.memHeightMap) != 0 && len(c.memHashMap) != 0 {
			return ErrMissingLatestBlock
		}
		strHash = string(block.GetHash())
		hexHeight = block.GetHeight().Text(16)
	} else {
		if bytes.Compare(c.latestBlock.GetHash(), block.GetPrevHash()) != 0 {
			return ErrWrongPrevHash
		}

		hexHeight = block.GetHeight().Text(16)

		if c.memHeightMap[hexHeight] != nil {
			return ErrHasSameHeightBlock
		}

		strHash = string(block.GetHash())
		if c.memHashMap[strHash] != nil {
			err = ErrHasSameHashBlock
		}
	}

	c.latestBlock = block.(*Block)
	c.memHashMap[strHash] = append(c.memHashMap[strHash], c.latestBlock)
	c.memHeightMap[hexHeight] = c.memHeightMap[hexHeight]
	return err
}

// Audit will look through the whole chain and check the integrity of chain
func (c *MemChain) Audit() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	current := *c.memHeightMap[c.latestBlock.GetHeight().Text(16)]
	if bytes.Compare(current.GetHash(), c.latestBlock.GetHash()) != 0 {
		return false
	}

	for {
		if current.GetPrevHash() == nil {
			if current.GetHeight() != big0 {
				return false
			}
			return true
		}

		currents := c.memHashMap[string(current.GetPrevHash())]
		if currents == nil {
			return false
		}

		for k := range currents {
			b := *currents[k]
			if b.GetHeight() == new(big.Int).Add(current.GetHeight(), big1) {
				current = b
				break
			}
		}

		var b = c.memHeightMap[current.GetHeight().Text(16)]
		if b == nil {
			return false
		}

		block := *b
		if bytes.Compare(current.GetHash(), block.GetHash()) != 0 {
			return false
		}
	}
}

func (c *MemChain) GetBlockByHeight(height *big.Int) abstract.Block {
	c.mu.RLock()
	defer c.mu.RUnlock()

	resultBlock := c.memHeightMap[height.Text(16)]
	return resultBlock
}

func (c *MemChain) GetBlockByHash(hash []byte) []abstract.Block {
	c.mu.RLock()
	defer c.mu.RUnlock()

	blocks := c.memHashMap[string(hash)]
	resultBlocks := make([]abstract.Block, len(blocks))
	for i, v := range c.memHashMap[string(hash)] {
		resultBlocks[i] = abstract.Block(v)
	}

	return resultBlocks
}

func NewMemChain() *MemChain {
	return &MemChain{
		latestBlock:  nil,
		mu:           &sync.RWMutex{},
		memHashMap:   make(map[string][]*Block),
		memHeightMap: make(map[string]*Block),
	}
}
