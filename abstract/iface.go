package abstract

import "math/big"

type Block interface {
	GetHeight() *big.Int
	GetHash() []byte
	GetPrevHash() []byte
	GetContent() []byte
}

type Chain interface {
	Push(block Block) error
	Audit() bool
	GetBlockByHeight(height *big.Int) Block
	GetBlockByHash(hash []byte) []Block
}
