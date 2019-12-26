package main

import "fmt"

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d Prev.Hash: %s Curr.Hash: %s Data: %s Timetamp: %s \n",
			block.Index, block.PreBlockHash, block.Hash, block.Data, block.Timestamp)
	}
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

func (bc *BlockChain) AppendBlock(new *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, new)
		return

	}
	if isValid(*new, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, new)
	}
}
func isValid(new, old Block) bool {
	if new.Index != old.Index+1 {
		return false
	}
	if new.PreBlockHash != old.Hash {
		return false
	}
	if calculateHash(new) != new.Hash {
		return false
	}

	return true
}
