package model

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	Number             uint64
	Hash               string
	ParentHash         string
	Sha3Uncles         string
	TransactionsRoot   string
	StateRoot          string
	ReceiptsRoot       string
	Miner              string
	Difficulty         uint64
	ExtraData          string
	GasLimit           uint64
	GasUsed            uint64
	Timestamp          uint64
}