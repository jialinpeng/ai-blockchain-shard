package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"math/big"
	"time"
)

// Transaction 交易结构
type Transaction struct {
	Sender      string   // 发送方地址
	Recipient   string   // 接收方地址
	Amount      *big.Int // 交易金额
	Nonce       uint64   // 交易序号
	Timestamp   time.Time // 时间戳
	Hash        []byte   // 交易哈希
	Relayed     bool     // 是否为中继交易
	FromShard   uint64   // 来源分片
	ToShard     uint64   // 目标分片
	Signature   []byte   // 交易签名
}

// NewTransaction 创建新交易
func NewTransaction(from, to string, amount *big.Int, nonce uint64, fromShard, toShard uint64) *Transaction {
	tx := &Transaction{
		Sender:    from,
		Recipient: to,
		Amount:    amount,
		Nonce:     nonce,
		Timestamp: time.Now(),
		Relayed:   false,
		FromShard: fromShard,
		ToShard:   toShard,
	}
	
	// 计算交易哈希
	hash := sha256.Sum256(tx.Encode())
	tx.Hash = hash[:]
	
	return tx
}

// Encode 交易编码
func (tx *Transaction) Encode() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(tx)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

// Decode 解码交易
func DecodeTransaction(data []byte) *Transaction {
	var tx Transaction
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&tx)
	if err != nil {
		panic(err)
	}
	return &tx
}

// IsCrossShard 判断是否为跨分片交易
func (tx *Transaction) IsCrossShard() bool {
	return tx.FromShard != tx.ToShard
}