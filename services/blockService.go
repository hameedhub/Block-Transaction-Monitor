package services

import (
	"ERC1155/model"
	"ERC1155/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/umbracle/go-web3/blocktracker"
	"github.com/umbracle/go-web3/jsonrpc"
	"gorm.io/gorm"
	"log"
	"time"
)

type BlocService interface {
	TrackBlock()
	GetTransaction(map[string]string, map[string]time.Time, int) ([]model.Transaction, error)
	SaveTransaction(log []byte, client *jsonrpc.Client)
}

type blockService struct {
	db *gorm.DB
	config util.Config
}

func NewBlockService(db *gorm.DB, config util.Config) BlocService {
	return blockService{
		db,config,
	}
}
func (s blockService) TrackBlock()  {
	client, err := jsonrpc.NewClient(s.config.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	tracker := blocktracker.NewBlockTracker(client.Eth())
	if err := tracker.Init(); err != nil {
		log.Fatal(err)
	}
	go tracker.Start()
	sub := tracker.Subscribe()
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case evnt:= <-sub:
				for _, log:= range evnt.Added{
					s.db.Create(&model.Block{Number: log.Number, Hash :log.Hash.String(),
						ParentHash : log.ParentHash.String(),
						Sha3Uncles : log.Sha3Uncles.String(),
						TransactionsRoot : log.TransactionsRoot.String(),
						StateRoot       : log.StateRoot.String(),
						ReceiptsRoot    : log.ReceiptsRoot.String(),
						Miner           : log.Miner.String(),
						Difficulty      : log.Difficulty.Uint64(),
						GasLimit        : log.GasLimit,
						GasUsed         : log.GasUsed,
						Timestamp       : log.Timestamp })
					rawLog, _ := json.Marshal(log)
					s.SaveTransaction(rawLog, client)
				}
			case <-ctx.Done():
			}
		}
	}()

}

func (s blockService) GetTransaction(query map[string]string, date map[string]time.Time, row int) (transactions []model.Transaction, err error)  {
	Rows := 100
	if row >0 {
		Rows = row
	}
	subQuery := s.db
	for i, v := range query {
		if v != "" {
			qy := fmt.Sprintf("\"%s\"  = '%v'", i, v)
			subQuery = subQuery.Where(qy)
		}
	}
	//subQuery = s.db.Debug()
	if date != nil {
		subQuery = subQuery.Where("created_at BETWEEN ? AND ?", date["from_date"], date["to_date"])
	}

	subQuery.Debug().Limit(Rows).Group("id").Order("created_at DESC").Find(&transactions)
	return
}

func (s blockService) SaveTransaction(log []byte, client *jsonrpc.Client) {

	rawString := string(log)
	rawBlock := &model.RawBlock{}
	json.Unmarshal([]byte(rawString), &rawBlock)

	for _, hash := range rawBlock.Transaction {
		var from string
		var to string
		var ethFrom float64
		var ethTo float64
		transaction := util.GetTransactionByHash(hash, client)

		if len(transaction.From.Bytes()) > 0 {
			from = transaction.From.String()
			ethFrom= util.GetBalance(from, client)
		}
		if len(transaction.To.Bytes()) >0{
			to = transaction.To.String()
			ethTo= util.GetBalance(to, client)
		}
			s.db.Create(&model.Transaction{
				BalFrom: ethFrom,
				BalTo: ethTo,
				Hash    : transaction.Hash.String(),
				From    : transaction.From.String(),
				To      : transaction.To.String(),
				GasPrice : transaction.GasPrice,
				Gas      : transaction.Gas,
				Value   : transaction.Value.String(),
				Nonce   : transaction.Nonce,
				BlockHash  : transaction.BlockHash.String(),
				BlockNumber : transaction.BlockNumber,
				TxnIndex   : transaction.TxnIndex,
				ChainID   : transaction.ChainID.String(),
				MaxPriorityFeePerGas : transaction.MaxPriorityFeePerGas.String(),
				MaxFeePerGas      : transaction.MaxFeePerGas.String(),
				CreatedAt: time.Now().UTC(),
			})
	}

}
