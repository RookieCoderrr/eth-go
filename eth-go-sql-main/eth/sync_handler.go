package eth

import (
	"context"
	"encoding/json"
	"fmt"
	. "git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/mq"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
	"time"
)

// 获取一个ETH区块数据
func GetBlock(client *ethclient.Client, number int64, isInsert bool) {
	blockNum := big.NewInt(number)
	block, td, txs, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		GlobalLogger.Error("block number: ", blockNum, "===>", err)
		return
	}
	GlobalLogger.Infof("Block: %v, TX count: %v", blockNum, len(block.Transactions()))
	// 保存交易数据
	_ = HandlerTransaction(client, block, txs, isInsert)
	// 保存区块数据
	if isInsert {
		nb := model.NewBlockFromEthTypes(block, td)
		_, err = BlockSvc.Create(nb)
		if err != nil {
			GlobalLogger.Errorf("BlockNumber %d Create Error", number)
		}
	}
}

func HandlerTransaction(client *ethclient.Client, block *types.Block, txs []ethclient.RPCTransaction, insert bool) error {
	txBulk := make([]*model.Transaction, 0)
	var logBulk []*model.Log
	var ttBulk []*model.TokenTransfer
	for _, tx := range txs {
		txHashHex := tx.GetTransactionObject().Hash().Hex()
		txInst, logBulkTemp, ttBulkTemp, err := GetTransaction(client, block, tx, insert)
		if err != nil {
			TxPushToRedis(block.Number().Int64(), txHashHex)
			fmt.Println("TxPushToRedis Err: ", err)
			return err
		}
		if txInst != nil {
			txBulk = append(txBulk, txInst)
		}
		logBulk = append(logBulk, logBulkTemp...)
		ttBulk = append(ttBulk, ttBulkTemp...)
	}

	fmt.Println("区块：", logBulk[0].BlockNumber, ",Logs数量：", len(logBulk), "Env: ", Env, "Insert: ", insert)
	if len(logBulk) > 0 {
		if insert {
			start := time.Now()
			_, _ = LogSvc.BatchCreate(logBulk, BulkSize)
			slice := time.Since(start)
			fmt.Println("区块：", logBulk[0].BlockNumber, ",Logs数量：", len(logBulk), ",总共用时：", slice)
		}
		// 追块由于必须保证log的插入顺序,放到mq中处理
		if Env == "prod" {
			fmt.Println("正要发送到Mq：", logBulk[0].BlockNumber, ",Logs数量：", len(logBulk))
			jsonb, err := json.Marshal(logBulk)
			if err != nil {
				fmt.Println("转换成json出错：", err)
				return err
			}
			mq.PublishTxLog(MqExchangeName, jsonb, false)
			fmt.Println("发送区块到Mq：", logBulk[0].BlockNumber, ",Logs数量：", len(logBulk))
		}
	}
	if len(txBulk) > 0 && insert {
		_, err := TransactionSvc.BatchCreate(txBulk)
		if err != nil {
			for _, tx := range txBulk {
				_, _ = TransactionSvc.Create(tx)
			}
		}
	}
	if len(ttBulk) > 0 && insert {
		_, _ = TokenTransferSvc.BatchCreate(ttBulk)
	}

	return nil
}

// 获取一个ETH区块交易数据
func GetTransaction(client *ethclient.Client, block *types.Block, tx ethclient.RPCTransaction, insert bool) (*model.Transaction, []*model.Log, []*model.TokenTransfer, error) {
	txHash := tx.GetTransactionObject().Hash()
	txHashHex := txHash.Hex()
	ethTxReceipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil && err != ethereum.NotFound {
		GlobalLogger.Error("client.TransactionReceipt error, hash value: ", txHashHex)
		GlobalLogger.Error(err)
		return nil, nil, nil, err
	}
	if ethTxReceipt == nil {
		//失败的交易直接返回
		return nil, nil, nil, err
	}
	txInst := model.NewTransactionFromEthTypes(block, tx.From.String(), tx.GetTransactionObject(), ethTxReceipt)
	if insert {
		go appendBalanceToChannel(client, txInst.FromAddress, block)
		var contractInst *model.Contract
		if txInst.ToAddress != "" {
			go appendBalanceToChannel(client, txInst.ToAddress, block)
		} else {
			contractInst = model.NewContractFromTransactionAttrs(txInst)
			ContractSvc.Create(contractInst)
		}
	}

	var logBulk []*model.Log
	var ttBulk []*model.TokenTransfer
	if len(ethTxReceipt.Logs) > 0 {
		logBulk, ttBulk = handleReceiptLogs(txInst, ethTxReceipt)
	}

	return txInst, logBulk, ttBulk, nil
}

// 获取一个ETH区块交易数据
func GetTransactionByTxHash(client *ethclient.Client, block *types.Block, txHashHex string) (*model.Transaction, []*model.Log, []*model.TokenTransfer, error) {
	txHash := common.HexToHash(txHashHex)
	ethTxReceipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil && err != ethereum.NotFound {
		GlobalLogger.Error("client.TransactionReceipt error, hash value: ", txHashHex)
		GlobalLogger.Error(err)
		return nil, nil, nil, err
	}
	if ethTxReceipt == nil {
		//失败的交易直接返回
		return nil, nil, nil, err
	}

	ethRtx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		GlobalLogger.Error("client.TransactionByHash error, hash value: ", txHashHex)
		GlobalLogger.Error(err)
		return nil, nil, nil, err
	}

	ethTx := ethRtx.GetTransactionObject()
	txInst := model.NewTransactionFromEthTypes(block, ethRtx.From.String(), ethTx, ethTxReceipt)

	go appendBalanceToChannel(client, txInst.FromAddress, block)

	var contractInst *model.Contract

	if txInst.ToAddress != "" {
		go appendBalanceToChannel(client, txInst.ToAddress, block)
	} else {
		contractInst = model.NewContractFromTransactionAttrs(txInst)
		ContractSvc.Create(contractInst)
	}
	var logBulk []*model.Log
	var ttBulk []*model.TokenTransfer
	if len(ethTxReceipt.Logs) > 0 {
		logBulk, ttBulk = handleReceiptLogs(txInst, ethTxReceipt)
	}

	return txInst, logBulk, ttBulk, nil
}

func appendBalanceToChannel(client *ethclient.Client, address string, block *types.Block) {
	return
	bValue := float64(0)
	// bValue, _ := getBalanceValue(client, address, block.Number())
	// cv := NewBalanceDataMap(address, bValue, block.Number().Int64())
	// balanceSaveChan <- *cv
	cacheKey := ChainName + "-balance-" + address
	if exist, _ := RedisRepo.IfExist(cacheKey); exist {
		return
	}
	inst, _ := BalanceSvc.QueryByAddress(address)
	if inst == nil {
		inst = model.NewBalanceFromFields(address, block.Number().Int64(), bValue)
		_, _ = BalanceSvc.Create(inst)
	}
	RedisRepo.Set(cacheKey, []byte{'1'}, 72*time.Hour)
}

func getBalanceValue(client *ethclient.Client, address string, blockNumber *big.Int) (float64, error) {
	blc, err := client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber)
	if err != nil {
		// GlobalLogger.Errorf("client.BalanceAt error, address: %v, blockNumber: %v", address, blockNumber.String())
		// GlobalLogger.Error(err)
		return 0, err
	}
	return float64(blc.Int64()), nil
}

func handleReceiptLogs(tx *model.Transaction, receipt *types.Receipt) ([]*model.Log, []*model.TokenTransfer) {
	var logBulk []*model.Log
	var ttBulk []*model.TokenTransfer
	logCount := len(receipt.Logs)
	for i := 0; i < logCount; i++ {
		lg := receipt.Logs[i]
		inst := model.NewLogFromEthTypes(lg)
		topic0 := GetLogTopicHash(lg, 0)
		if topic0 != "" && (topic0 == TransferEventTopic || topic0 == Erc1155TransferSingleEventTopic || topic0 == Erc1155TransferBatchEventTopic) {
			tt := HandleTokenTransfer(tx, lg)
			if tt != nil {
				ttBulk = append(ttBulk, tt)
			}
		}
		logBulk = append(logBulk, inst)
	}
	return logBulk, ttBulk
}

func GetLogTopicHash(lg *types.Log, index int) string {
	topicHash := ""
	if len(lg.Topics) > 0 {
		topicHash = lg.Topics[index].Hex()
	}
	return topicHash
}

func HandleTokenTransfer(tx *model.Transaction, lg *types.Log) (tt *model.TokenTransfer) {
	defer func() {
		if err := recover(); err != nil {
			m, _ := json.Marshal(lg)
			GlobalLogger.Error("HandleTokenTransfer Error: ", err)
			GlobalLogger.Error("HandleTokenTransfer block: ", tx.BlockNumber)
			GlobalLogger.Error("HandleTokenTransfer tx: ", tx.TransactionHash)
			GlobalLogger.Error("HandleTokenTransfer log: ", lg.Index)
			GlobalLogger.Error("HandleTokenTransfer Data: ")
			GlobalLogger.Error(string(m))
			tt = nil
		}
	}()

	topic := GetLogTopicHash(lg, 0)
	value := new(big.Int)
	eventName := ""
	fromAddr := ""
	toAddr := ""
	if topic == TransferEventTopic {
		eventName = "Transfer"
		topicCount := len(lg.Topics)
		_hex := common.Bytes2Hex(lg.Data)
		hexStr := utils.TrimLeftZeroes(_hex)
		if topicCount == 1 {
			fromAddr = common.HexToAddress(hexStr[0:40]).String()
			hexStr = hexStr[40:]
			hexStr = utils.TrimLeftZeroes(hexStr)
			if len(hexStr) > 40 {
				toAddr = common.HexToAddress(hexStr[0:40]).String()
				hexStr = hexStr[40:]
				hexStr = utils.TrimLeftZeroes(hexStr)
			} else {
				toAddr = fromAddr
			}
			value = utils.HexToBigInt(hexStr)
		} else if topicCount == 2 {
			fromAddr = GetLogTopicHash(lg, 1)
			toAddr = common.HexToAddress(hexStr[0:40]).String()
			hexStr = hexStr[40:]
			hexStr = utils.TrimLeftZeroes(hexStr)
			value = utils.HexToBigInt(hexStr)
		} else {
			fromAddr = GetLogTopicHash(lg, 1)
			toAddr = GetLogTopicHash(lg, 2)
			if len(hexStr) > 2 {
				value = utils.HexToBigInt(hexStr)
			}
		}
	} else if topic == Erc1155TransferSingleEventTopic {
		eventName = "TransferSingle"
	} else if topic == Erc1155TransferBatchEventTopic {
		eventName = "TransferBatch"
	}
	tt = model.NewTokenTransferFromFields(tx, lg, topic, eventName, fromAddr, toAddr, value)
	return tt
}

// 重跑次数加入到map中，超过3次依然没有成功的直接放弃
var failureTxCounter sync.Map

func TxPushToRedis(blockNumber int64, txHashHex string) {
	// 出错的txHash放入redis, 由专职协程处理,  值格式: 区块号!-_-!交易hash值
	value := fmt.Sprintf("%v!-_-!%v", blockNumber, txHashHex)
	if counter, exist := failureTxCounter.Load(value); exist {
		counterInt := counter.(int)
		if counterInt > 3 {
			failureTxCounter.Delete(value)
			return
		}
		failureTxCounter.Store(value, counterInt+1)
	} else {
		failureTxCounter.Store(value, 1)
	}
	length, err := RedisRepo.LPush(GlobalTxFailedQueue, value)
	if err != nil {
		GlobalLogger.Error("RedisRepo.LPush error, hash value: ", value)
		GlobalLogger.Error(err)
	}
	if length > 0 && length%100 == 0 {
		GlobalLogger.Infof("Redis Query Size:%d", length)
	}
}
