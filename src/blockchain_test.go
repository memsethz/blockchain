package blockchain

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more to Ivan")
	iterator := bc.Iterator()
	for {
		block := iterator.Next()
		if block != nil {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			pow := NewProofOfWork(block)
			fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
		} else {
			break
		}
	}
}
