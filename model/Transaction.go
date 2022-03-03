package model

import "time"

type Transaction struct {
	ID    uint `gorm:"primarykey" json:"id"`
	BalFrom float64 	`json:"balance_sender"`
	BalTo 	float64  	`json:"balance receiver"`
	Hash     string		`json:"hash"`
	From     string 	`json:"sender"`
	To       string		`json:"receiver"`
	Input    string		`json:"-"`
	GasPrice uint64		`json:"gas_price"`
	Gas      uint64		`json:"gas"`
	Value   string		`json:"value"`
	Nonce    uint64		`json:"nonce"`
	BlockHash   string	`json:"block_hash"`
	BlockNumber uint64	`json:"block_number"`
	TxnIndex    uint64	`json:"txn_index"`	
	ChainID    string		`json:"chain_id"`
	MaxPriorityFeePerGas string	`json:"max_priority_fee_per_gas"`
	MaxFeePerGas         string	`json:"max_fee_per_gas"`
	CreatedAt		time.Time `json:"created_at"`
}
