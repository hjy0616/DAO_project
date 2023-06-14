package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
	Creator   string
}

type Stakeholder struct {
	Address string
	Stake   int
}

type Blockchain struct {
	Chain        []Block
	Stakeholders []Stakeholder
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createBlock(oldBlock Block, data string, creator string) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Creator = creator

	return newBlock
}

func main() {
	var bc Blockchain

	genesisBlock := Block{}
	genesisBlock = Block{0, time.Now().String(), "", calculateHash(genesisBlock), "", ""}
	bc.Chain = append(bc.Chain, genesisBlock)

	bc.Stakeholders = append(bc.Stakeholders, Stakeholder{"address1", 100})
	bc.Stakeholders = append(bc.Stakeholders, Stakeholder{"address2", 200})

	for i := 0; i < 10; i++ {
		blockData := "Block " + strconv.Itoa(i) + " Data"

		totalStake := 0
		for _, s := range bc.Stakeholders {
			totalStake += s.Stake
		}

		r := rand.Intn(totalStake)

		sum := 0
		creator := ""
		for _, s := range bc.Stakeholders {
			sum += s.Stake
			if r < sum {
				creator = s.Address
				break
			}
		}

		bc.Chain = append(bc.Chain, createBlock(bc.Chain[i], blockData, creator))
	}
}
