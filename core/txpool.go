package core

import (
	"sync"
)

// TxPool 交易池
type TxPool struct {
	PendingTxs   []*Transaction     // 待处理交易
	RelayTxs     []*Transaction     // 中继交易
	mutex        sync.Mutex
}

// NewTxPool 创建新的交易池
func NewTxPool() *TxPool {
	return &TxPool{
		PendingTxs: make([]*Transaction, 0),
		RelayTxs:   make([]*Transaction, 0),
	}
}

// AddTransaction 添加交易到交易池
func (tp *TxPool) AddTransaction(tx *Transaction) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	
	if tx.IsCrossShard() {
		tp.RelayTxs = append(tp.RelayTxs, tx)
	} else {
		tp.PendingTxs = append(tp.PendingTxs, tx)
	}
}

// AddTransactions 批量添加交易
func (tp *TxPool) AddTransactions(txs []*Transaction) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	
	for _, tx := range txs {
		if tx.IsCrossShard() {
			tp.RelayTxs = append(tp.RelayTxs, tx)
		} else {
			tp.PendingTxs = append(tp.PendingTxs, tx)
		}
	}
}

// GetPendingTransactions 获取待处理交易
func (tp *TxPool) GetPendingTransactions(limit int) []*Transaction {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	
	end := limit
	if end > len(tp.PendingTxs) {
		end = len(tp.PendingTxs)
	}
	
	txs := tp.PendingTxs[:end]
	tp.PendingTxs = tp.PendingTxs[end:]
	
	return txs
}

// GetRelayTransactions 获取中继交易
func (tp *TxPool) GetRelayTransactions(limit int) []*Transaction {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	
	end := limit
	if end > len(tp.RelayTxs) {
		end = len(tp.RelayTxs)
	}
	
	txs := tp.RelayTxs[:end]
	tp.RelayTxs = tp.RelayTxs[end:]
	
	return txs
}

// PendingCount 获取待处理交易数量
func (tp *TxPool) PendingCount() int {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	return len(tp.PendingTxs)
}

// RelayCount 获取中继交易数量
func (tp *TxPool) RelayCount() int {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()
	return len(tp.RelayTxs)
}