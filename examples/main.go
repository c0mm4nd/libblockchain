package main

import (
	"encoding/json"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/maoxs2/libblockchain"
)

var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	big0     = big.NewInt(0)
	big1     = big.NewInt(1)
	big10000 = big.NewInt(10000)
)

func main() {
	chain := libblockchain.NewMemChain()

	height := big0
	genesisBlock := libblockchain.NewGenesisBlock(time.Now().UnixNano(), randomItem())
	hash := genesisBlock.GetHash()

	err := chain.Push(genesisBlock)
	if err != nil {
		log.Println(err)
	}

	for {
		height.Add(height, big1)
		block := libblockchain.NewBlock(height, hash, time.Now().UnixNano(), randomItem())
		err := chain.Push(block)
		if err != nil {
			log.Println(err)
		} else {
			b, _ := json.Marshal(block)
			log.Println("added block", string(b))
		}

		if new(big.Int).Mod(height, big10000) == big0 {
			if !chain.Audit() {
				return
			}
		}

		hash = block.GetHash()
	}

}

func randStr(l int) string {
	b := make([]byte, l)

	for i := range b {
		l := len(letterBytes)
		b[i] = letterBytes[rand.Intn(l)]
	}

	return string(b)
}

func randomItem() []byte {
	item := map[string]interface{}{
		randStr(10): randStr(100),
		randStr(10): randStr(100),
		randStr(10): randStr(100),
	}

	b, _ := json.Marshal(item)
	return b
}
