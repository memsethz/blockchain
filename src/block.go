package blockchain

import (
	"time"
)

type Block struct {
	Timestamp     int64 //current block creatdTime
	Data          []byte
	PrevBlockHash []byte //prev block hash
	Hash          []byte //current block hash
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//计算hash
//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{ b.PrevBlockHash, b.Data, timestamp, }, []byte{})
//	hash := sha256.Sum256(headers)
//
//	b.Hash = hash[:]
//}
