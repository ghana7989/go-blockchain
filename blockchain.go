package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.previousHash = previousHash
	b.nonce = nonce
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp			%d\n", b.timestamp)
	fmt.Printf("nonce				%d\n", b.nonce)
	fmt.Printf("previousHash			%x\n", b.previousHash)
	fmt.Printf("transactions			%s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Timestamp    int64    `json:"timestamp"`
		Transactions []string `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func main() {
	blockchain := NewBlockchain()
	prevHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, prevHash)
	prevHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(4, prevHash)
	blockchain.Print()
}
