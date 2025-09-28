package shard

import (
	"ai-blockchain-shard/core"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Node 分片节点
type Node struct {
	ID         uint64
	ShardID    uint64
	Address    string
	Blockchain *core.BlockChain
}

// NewNode 创建新节点
func NewNode(id, shardID uint64, address string) *Node {
	return &Node{
		ID:         id,
		ShardID:    shardID,
		Address:    address,
		Blockchain: core.NewBlockChain(shardID),
	}
}

// GenerateSampleTransactions 生成示例交易
func (n *Node) GenerateSampleTransactions(count int) []*core.Transaction {
	var transactions []*core.Transaction
	
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < count; i++ {
		// 随机生成发送方和接收方
		from := fmt.Sprintf("account_%d", rand.Intn(1000))
		to := fmt.Sprintf("account_%d", rand.Intn(1000))
		
		// 随机生成金额
		amount := big.NewInt(rand.Int63n(1000) + 1)
		
		// 随机生成nonce
		nonce := uint64(rand.Int63n(100))
		
		// 确定分片
		fromShard := n.ShardID
		toShard := n.ShardID
		
		// 有一定概率生成跨分片交易
		if rand.Float32() < 0.3 { // 30%概率生成跨分片交易
			toShard = (toShard + 1) % 4 // 假设有4个分片
		}
		
		tx := core.NewTransaction(from, to, amount, nonce, fromShard, toShard)
		transactions = append(transactions, tx)
	}
	
	return transactions
}

// AddTransactions 添加交易到节点的区块链交易池
func (n *Node) AddTransactions(txs []*core.Transaction) {
	n.Blockchain.TxPool.AddTransactions(txs)
}

// MineBlock 挖掘新区块
func (n *Node) MineBlock() *core.Block {
	fmt.Printf("Node %d in Shard %d is mining a new block...\n", n.ID, n.ShardID)
	
	block := n.Blockchain.GenerateBlock()
	
	if n.Blockchain.AddBlock(block) {
		fmt.Printf("Node %d successfully mined block #%d\n", n.ID, block.Header.Number)
		return block
	} else {
		fmt.Printf("Node %d failed to mine block #%d\n", n.ID, block.Header.Number)
		return nil
	}
}

// PrintStatus 打印节点状态
func (n *Node) PrintStatus() {
	fmt.Printf("Node ID: %d\n", n.ID)
	fmt.Printf("Shard ID: %d\n", n.ShardID)
	fmt.Printf("Address: %s\n", n.Address)
	fmt.Printf("Current blockchain height: %d\n", n.Blockchain.CurrentHeight())
	fmt.Printf("Pending transactions: %d\n", n.Blockchain.TxPool.PendingCount())
	fmt.Printf("Relay transactions: %d\n", n.Blockchain.TxPool.RelayCount())
	fmt.Println("---")
}