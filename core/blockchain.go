package core

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// BlockChain 区块链结构
type BlockChain struct {
	Blocks     []*Block
	ShardID    uint64
	TxPool     *TxPool
}

// NewBlockChain 创建新的区块链
func NewBlockChain(shardID uint64) *BlockChain {
	genesisBlock := createGenesisBlock(shardID)
	
	bc := &BlockChain{
		Blocks:  []*Block{genesisBlock},
		ShardID: shardID,
		TxPool:  NewTxPool(),
	}
	
	return bc
}

// createGenesisBlock 创建创世区块
func createGenesisBlock(shardID uint64) *Block {
	header := &BlockHeader{
		ParentBlockHash: []byte{}, // 创世区块没有父区块
		StateRoot:       []byte{}, // 初始状态根
		TxRoot:          []byte{}, // 没有交易
		Number:          0,        // 区块高度为0
		Time:            time.Now(),
		ShardID:         shardID,
	}
	
	// 计算区块头哈希
	headerHash := header.Hash()
	
	block := &Block{
		Header: header,
		Body:   []*Transaction{}, // 没有交易
		Hash:   headerHash,
	}
	
	return block
}

// AddBlock 添加区块到区块链
func (bc *BlockChain) AddBlock(block *Block) bool {
	// 验证区块
	if !bc.isValidBlock(block) {
		return false
	}
	
	// 添加到区块链
	bc.Blocks = append(bc.Blocks, block)
	return true
}

// isValidBlock 验证区块是否有效
func (bc *BlockChain) isValidBlock(block *Block) bool {
	// 检查区块高度
	if block.Header.Number != bc.CurrentHeight()+1 {
		return false
	}
	
	// 检查父区块哈希
	lastBlock := bc.LastBlock()
	if !bytes.Equal(block.Header.ParentBlockHash, lastBlock.Hash) {
		return false
	}
	
	// 检查交易根
	calculatedTxRoot := bc.calculateTxRoot(block.Body)
	if !bytes.Equal(block.Header.TxRoot, calculatedTxRoot) {
		return false
	}
	
	return true
}

// calculateTxRoot 计算交易根
func (bc *BlockChain) calculateTxRoot(transactions []*Transaction) []byte {
	if len(transactions) == 0 {
		return []byte{}
	}
	
	// 简化实现：将所有交易哈希连接后计算哈希
	var txHashes []byte
	for _, tx := range transactions {
		txHashes = append(txHashes, tx.Hash...)
	}
	
	hash := sha256.Sum256(txHashes)
	return hash[:]
}

// LastBlock 获取最后一个区块
func (bc *BlockChain) LastBlock() *Block {
	if len(bc.Blocks) == 0 {
		return nil
	}
	return bc.Blocks[len(bc.Blocks)-1]
}

// CurrentHeight 获取当前区块高度
func (bc *BlockChain) CurrentHeight() uint64 {
	lastBlock := bc.LastBlock()
	if lastBlock == nil {
		return 0
	}
	return lastBlock.Header.Number
}

// GenerateBlock 生成新区块
func (bc *BlockChain) GenerateBlock() *Block {
	// 从交易池获取交易
	transactions := bc.TxPool.GetPendingTransactions(100) // 限制区块大小为100个交易
	
	// 创建区块头
	header := &BlockHeader{
		ParentBlockHash: bc.LastBlock().Hash,
		StateRoot:       []byte{}, // 简化实现，不处理状态
		TxRoot:          bc.calculateTxRoot(transactions),
		Number:          bc.CurrentHeight() + 1,
		Time:            time.Now(),
		ShardID:         bc.ShardID,
	}
	
	// 创建区块
	block := NewBlock(header, transactions)
	return block
}