package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64  //区块编号
	Timestamp    int64  //区块时间戳
	PreBlockHash string //上一个区块hash值
	Hash         string //当前区块hash

	Data string //区块数据
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PreBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])

}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.Data = data
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//创世区块
func GenerateGenesisBlock() Block {
	preBlock := Block{
		Index: -1,
		Hash:  "",
	}
	return GenerateNewBlock(preBlock, "Genesis Block")
}
