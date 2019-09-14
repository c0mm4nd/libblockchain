# libblockchain

Make blockchain integration easy

## Intro

This blockchain has no consensus, which means that you can integrate it into any application and act as any role like db, mq etc.

This structure is thread(routine?)-safe and can avoid the covering block from the hash collide

## usage 

look into the example in `example` folder. Or check the [godoc.org](https://godoc.org/github.com/maoxs2/libblockchain)

## TODO

- storageChain, which automatically saves the block and the chain into the backend key-value database
- optimize the memChain
