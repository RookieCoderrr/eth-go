package test

import (
	"context"
	"encoding/json"
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"testing"
)

// var url ="http://119.28.50.210:8545" //matic
//var url="http://119.28.131.74:8545"
var url="ws://101.32.199.11:8546" //bsc
// 测试同步
func TestGetLastBlockInfo(t *testing.T) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	progress,_ := client.SyncProgress(context.Background())
	proJson,_ := json.Marshal(progress)
	fmt.Println("Progress:", string(proJson))

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Header BlockNumber:",header.Number.String())
}

// 测试获取一个块的所有关联数据
func TestGetBlockData(t *testing.T) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	blockNumber := big.NewInt(10009378)
	block,_, _, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	blockContent, _ := json.MarshalIndent(block.Header(), "", "    ")
	fmt.Println("resp-block:\n", string(blockContent))
	fmt.Println("block-tx-count: ", len(block.Transactions()))

	for _, tt := range block.Transactions() {
		ethRtx, _, err := client.TransactionByHash(context.Background(), tt.Hash())
		if err != nil {
			fmt.Println("ethRtx error ======> ", err)
			continue
		}

		ethTxReceipt, err := client.TransactionReceipt(context.Background(), tt.Hash())
		if err != nil {
			fmt.Println("ethTxReceipt error ======> ", err)
			continue
		}

		ethTx := ethRtx.GetTransactionObject()
		txInst := model.NewTransactionFromEthTypes(block, ethRtx.From.String(), ethTx, ethTxReceipt)

		//fmt.Println("len(ethTxReceipt.Logs) ======> ", len(ethTxReceipt.Logs))
		if len(ethTxReceipt.Logs) > 0 {
			var logBulk []*model.Log
			var ttBulk []*model.TokenTransfer
			logCount := len(ethTxReceipt.Logs)
			for i := 0; i < logCount; i++ {
				lg := ethTxReceipt.Logs[i]
				fmt.Printf("%d - %s - %d - %d \n", lg.BlockNumber,lg.TxHash,lg.TxIndex,lg.Index)
				inst := model.NewLogFromEthTypes(lg)
				topic0 := eth.GetLogTopicHash(lg, 0)
				if topic0 != "" && (topic0 == eth.TransferEventTopic || topic0 == eth.Erc1155TransferSingleEventTopic || topic0 == eth.Erc1155TransferBatchEventTopic) {
					tt := eth.HandleTokenTransfer(txInst, lg)
					if tt != nil {
						ttBulk = append(ttBulk, tt)
					}
				}

				logBulk = append(logBulk, inst)
			}
		}
	}
}

func TestTransaction(t *testing.T){
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/24d60bb2332948ada9a07828f0923a12")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	receipt,err := client.TransactionReceipt(context.Background(), common.HexToHash("0xbab68d1c70ae2c414cab0375e1d36092a5334a56313194f403987858ea4f3fe7"))
	if err!=nil{
		t.Error(err)
	}
	tx,isPanding,err := client.TransactionByHash(context.Background(),common.HexToHash("0xbab68d1c70ae2c414cab0375e1d36092a5334a56313194f403987858ea4f3fe7"))
	if err!=nil{
		t.Error(err)
	}
	t.Log(receipt)
	t.Log(utils.HexToBigInt(*tx.BlockNumber))
	t.Log(isPanding)
}

type rpcBlock struct {
	Hash            common.Hash      `json:"hash"`
	Transactions    []RPCTransaction `json:"transactions"`
	UncleHashes     []common.Hash    `json:"uncles"`
	TotalDifficulty *hexutil.Big  `json:"totalDifficulty"`
}
type RPCTransaction struct {
	// caches
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func TestTime(t *testing.T){
}


