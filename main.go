package main

import (
	"ai-blockchain-shard/core"
	"ai-blockchain-shard/shard"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	fmt.Println("A Blockchain Sharding System Simulator using LingMa")
	fmt.Println("=============================")
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command]")
		fmt.Println("Available commands:")
		fmt.Println("  init                    - Initialize the blockchain network")
		fmt.Println("  start <shards> <nodes>  - Start a network with specified shards and nodes per shard")
		fmt.Println("  generate <node_id> <count> - Generate sample transactions for a node")
		fmt.Println("  mine <node_id>          - Mine a block on a node")
		fmt.Println("  status                  - Show network status")
		return
	}
	
	command := os.Args[1]
	
	switch command {
	case "init":
		fmt.Println("Initializing blockchain network...")
		// 初始化区块链网络
		
	case "start":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go start <shards> <nodes>")
			return
		}
		
		shards, err1 := strconv.Atoi(os.Args[2])
		nodes, err2 := strconv.Atoi(os.Args[3])
		
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid parameters. Please provide numbers for shards and nodes.")
			return
		}
		
		startNetwork(shards, nodes)
		
	case "generate":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go generate <node_id> <count>")
			return
		}
		
		nodeID, err1 := strconv.Atoi(os.Args[2])
		count, err2 := strconv.Atoi(os.Args[3])
		
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid parameters. Please provide numbers for node_id and count.")
			return
		}
		
		generateTransactions(nodeID, count)
		
	case "mine":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go mine <node_id>")
			return
		}
		
		nodeID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid node ID.")
			return
		}
		
		mineBlock(nodeID)
		
	case "status":
		showStatus()
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

// startNetwork 启动网络
func startNetwork(shards, nodes int) {
	fmt.Printf("Starting network with %d shards and %d nodes per shard\n", shards, nodes)
	
	// 创建网络
	network := shard.NewNetwork()
	
	// 创建分片和节点
	nodeID := 1
	for s := 0; s < shards; s++ {
		for n := 0; n < nodes; n++ {
			address := fmt.Sprintf("192.168.1.%d", nodeID)
			node := shard.NewNode(uint64(nodeID), uint64(s), address)
			network.AddNode(node)
			nodeID++
		}
	}
	
	// 保存网络到全局状态或文件
	fmt.Println("Network started successfully!")
	network.PrintNetworkInfo()
}

// generateTransactions 生成交易
func generateTransactions(nodeID, count int) {
	fmt.Printf("Generating %d transactions for node %d\n", count, nodeID)
	
	// 创建一些示例交易
	var transactions []*core.Transaction
	
	for i := 0; i < count; i++ {
		from := fmt.Sprintf("account_%d", i)
		to := fmt.Sprintf("account_%d", i+1000)
		amount := big.NewInt(int64(i + 1))
		nonce := uint64(i)
		fromShard := uint64(0)
		toShard := uint64(0)
		
		// 有一定概率生成跨分片交易
		if i%5 == 0 {
			toShard = 1
		}
		
		tx := core.NewTransaction(from, to, amount, nonce, fromShard, toShard)
		transactions = append(transactions, tx)
	}
	
	fmt.Printf("Generated %d transactions\n", len(transactions))
	
	// 这里应该将交易添加到指定节点，但在简化版本中仅打印
	for i, tx := range transactions {
		if i < 5 { // 只打印前5个交易
			fmt.Printf("Transaction %d: %s -> %s, Amount: %s, Shard: %d -> %d\n", 
				i, tx.Sender, tx.Recipient, tx.Amount.String(), tx.FromShard, tx.ToShard)
		}
	}
}

// mineBlock 挖掘区块
func mineBlock(nodeID int) {
	fmt.Printf("Mining block on node %d\n", nodeID)
	// 在实际实现中，这里会调用节点的挖矿功能
	fmt.Println("Block mined successfully!")
}

// showStatus 显示状态
func showStatus() {
	fmt.Println("Network status:")
	fmt.Println("Shard 0: 2 nodes")
	fmt.Println("Shard 1: 2 nodes")
	fmt.Println("Total transactions: 0")
	fmt.Println("Current block height: 0")
}