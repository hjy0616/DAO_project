package main

import (
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

type Block struct {
	Transactions []string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{},
	}
}

func (b *Blockchain) AddBlock(transactions []string) {
	block := &Block{Transactions: transactions}
	b.Blocks = append(b.Blocks, block)
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock([]string{"tx1", "tx2"})

	db, err := bolt.Open("blockchain.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("BlocksBucket"))
		if err != nil {
			return err
		}

		gob.Register(Block{})
		var buf []byte
		enc := gob.NewEncoder(os.Stdout)
		for _, block := range bc.Blocks {
			err := enc.Encode(block)
			if err != nil {
				log.Fatal("encode error:", err)
			}
			buf = append(buf, os.Stdout.Bytes()...)
			os.Stdout.Reset()
		}
		err = b.Put([]byte("Blockchain"), buf)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

// 디코드 작업 코드
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	return *block
}

//(디코딩 과정) gob.NewDecoder 함수
