package pkg

import (
	"github.com/WSChong33/go_blockchain/internal"
)

type Blockchain struct {
	Blocks []*internal.Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*internal.Block{internal.NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := internal.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
