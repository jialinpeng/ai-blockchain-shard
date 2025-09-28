package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// BlockHeader 区块头定义
type BlockHeader struct {
	ParentBlockHash []byte    // 父区块哈希
	StateRoot       []byte    // 状态树根
	TxRoot          []byte    // 交易树根
	Number          uint64    // 区块高度
	Time            time.Time // 时间戳
	ShardID         uint64    // 分片ID
}

// Block 区块定义
type Block struct {
	Header    *BlockHeader
	Body      []*Transaction
	Hash      []byte
	Signature []byte // 区块签名
}

// NewBlock 创建新区块
func NewBlock(header *BlockHeader, body []*Transaction) *Block {
	block := &Block{
		Header: header,
		Body:   body,
	}
	
	// 计算区块哈希
	hash := sha256.Sum256(block.Encode())
	block.Hash = hash[:]
	
	return block
}

// Encode 区块编码
func (b *Block) Encode() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(b)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

// Decode 解码区块
func Decode(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return &block
}

// Hash 计算区块头哈希
func (bh *BlockHeader) Hash() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(bh)
	if err != nil {
		panic(err)
	}
	
	hash := sha256.Sum256(buff.Bytes())
	return hash[:]
}