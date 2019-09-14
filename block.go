package libblockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

type Block struct {
	Height   *big.Int
	Hash     []byte
	PrevHash []byte
	Content  []byte
}

func (b *Block) GetHash() []byte {
	return b.Hash
}

func (b *Block) GetPrevHash() []byte {
	return b.PrevHash
}

func (b *Block) GetContent() []byte {
	return b.Content
}

func (b *Block) GetHeight() *big.Int {
	return b.Height
}

func NewBlock(height *big.Int, prevHash []byte, timestamp int64, content []byte) *Block {
	// due to big follows big-endian, so deal all number in this endian
	bTimeStamp := make([]byte, 8)
	binary.BigEndian.PutUint64(bTimeStamp, uint64(timestamp))

	fullText := bytes.Join([][]byte{height.Bytes(), prevHash, bTimeStamp, content}, nil)
	hash := sha256.Sum256(fullText)

	return &Block{
		Height:   height,
		Hash:     hash[:],
		PrevHash: prevHash,
		Content:  content,
	}
}

func NewGenesisBlock(timestamp int64, content []byte) *Block {
	// due to big follows big-endian, so deal all number in this endian
	bTimeStamp := make([]byte, 8)
	binary.BigEndian.PutUint64(bTimeStamp[0:], uint64(timestamp))

	fullText := bytes.Join([][]byte{big0.Bytes(), nil, bTimeStamp[:], content}, nil)
	hash := sha256.Sum256(fullText)

	return &Block{
		Height:   big0,
		Hash:     hash[:],
		PrevHash: nil,
		Content:  content,
	}
}
