package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	fmt.Println("Genesis block has been created!")
	fmt.Println("Index:", Blockchain[0].Index)
	fmt.Println("Timestamp:", Blockchain[0].Timestamp)
	fmt.Println("BPM:", Blockchain[0].BPM)
	fmt.Println("Hash:", Blockchain[0].Hash)
	fmt.Println("PrevHash:", Blockchain[0].PrevHash)
}
