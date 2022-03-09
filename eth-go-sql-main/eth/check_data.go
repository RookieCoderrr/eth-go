package eth

import (
	"context"
	. "git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"strings"
)

func CheckTxAndLogData(client *ethclient.Client, number int64) {
	GlobalLogger.Infof("Block: %v", number)
	if !BlockSvc.NumberExist(number) {
		// 没有块数据，直接补全块
		GetBlock(client, number, true)
		return
	}
	//blockNum := big.NewInt(number)
	//block, _,txs, err := client.BlockByNumber(context.Background(), blockNum)
	//if err != nil {
	//	GlobalLogger.Error("block number: ", blockNum, "===>", err)
	//	return
	//}
	//
	//rpcTxCount := len(block.Transactions())
	//GlobalLogger.Infoln("BlockNumber:", blockNum, "TxCount:", rpcTxCount)
	//if rpcTxCount == 0 {
	//	// 调用接口返回的交易数量为 0 的不作处理
	//	return
	//}
	//
	//txListExist,_ := txRepo.QueryByBlockNumber(number)
	//txHashSet := make(map[string]*model.Transaction)
	//for _,tx:=range txListExist {
	//	txHashSet[tx.TransactionHash]=tx
	//}
	//
	//logListExist,_ := logRepo.QueryByBlockNumber(number)
	//logSet := make(map[int]bool)
	//for _,lg := range logListExist{
	//	logSet[lg.LogIndex] = true
	//}
	//
	//// 逐个检查tx中的Logs数量
	//txBulk := make([]*model.Transaction, 0)
	//var logBulk []*model.Log
	//var ttBulk []*model.TokenTransfer
	//for _, tx := range txs {
	//	txHash := tx.GetTransactionObject().Hash()
	//	txHashHex := txHash.Hex()
	//	ethTxReceipt, err := client.TransactionReceipt(context.Background(),txHash)
	//	if err != nil || ethTxReceipt==nil{
	//		GlobalLogger.Error("client.TransactionReceipt error, hash value: ", txHash)
	//		GlobalLogger.Error(err)
	//		TxPushToRedis(number, txHashHex)
	//		continue
	//	}
	//
	//	var txInst *model.Transaction
	//	if dbTx,exist := txHashSet[txHashHex];exist{
	//		txInst = dbTx
	//	}else{
	//		txInst = model.NewTransactionFromEthTypes(block, tx.From.String(), tx.GetTransactionObject(), ethTxReceipt)
	//		txBulk = append(txBulk, txInst)
	//	}
	//	if len(ethTxReceipt.Logs) > 0 {
	//		logBulkTemp, ttBulkTemp := handleMissingReceiptLogs(txInst, ethTxReceipt, logSet)
	//		logBulk = append(logBulk, logBulkTemp...)
	//		ttBulk = append(ttBulk, ttBulkTemp...)
	//	}
	//}
	//if len(logBulk) > 0 {
	//	_, _ = LogSvc.BatchCreate(logBulk,BulkSize)
	//}
	//if len(ttBulk) > 0 {
	//	_, _ = TokenTransferSvc.BatchCreate(ttBulk)
	//}
	//if len(txBulk) > 0 {
	//	_, err := TransactionSvc.BatchCreate(txBulk)
	//	if err!=nil{
	//		for _,tx := range txBulk{
	//			_,_ = TransactionSvc.Create(tx)
	//		}
	//	}
	//}
}

func handleMissingReceiptLogs(tx *model.Transaction, receipt *types.Receipt, logSet map[int]bool) ([]*model.Log, []*model.TokenTransfer) {
	var logBulk []*model.Log
	var ttBulk []*model.TokenTransfer
	logCount := len(receipt.Logs)
	for i := 0; i < logCount; i++ {
		lg := receipt.Logs[i]
		if logSet[int(lg.Index)] {
			continue
		}
		topic0 := GetLogTopicHash(lg, 0)
		if topic0 == "" || (topic0 != TransferEventTopic && topic0 != Erc1155TransferSingleEventTopic && topic0 != Erc1155TransferBatchEventTopic) {
			continue
		}
		conditions := map[string]interface{}{
			"block_number": tx.BlockNumber,
			"log_index":    lg.Index,
		}
		result, err := LogSvc.Query(conditions)
		if err != nil {
			GlobalLogger.Error("LogSvc.Query(conditions) error: ", conditions)
			continue
		}
		if result == nil || len(result) == 0 {
			inst := model.NewLogFromEthTypes(lg)
			tt := HandleTokenTransfer(tx, lg)
			if tt != nil {
				ttBulk = append(ttBulk, tt)
			}
			GlobalLogger.Infof("Fix Log BlockNumber:%d, TxHash:%s, LogIndex:%d\n", lg.BlockNumber, lg.TxHash, lg.Index)
			logBulk = append(logBulk, inst)
		}
	}
	return logBulk, ttBulk
}

// value格式: 区块号!-_-!交易hash值
func HandlerSingleTransactionForChecking(client *ethclient.Client, cacheValue string) {
	s := strings.Split(cacheValue, "!-_-!")
	number, _ := strconv.Atoi(s[0])
	txHashHex := s[1]

	blockNum := big.NewInt(int64(number))
	block, _, _, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		GlobalLogger.Error("block number: ", blockNum, "===>", err)
		return
	}

	txInst, logBulk, ttBulk, err := GetTransactionByTxHash(client, block, txHashHex)
	if err != nil {
		TxPushToRedis(int64(number), txHashHex)
		return
	}
	if txInst != nil {
		_, _ = TransactionSvc.Create(txInst)
	}
	if len(logBulk) > 0 {
		_, _ = LogSvc.BatchCreate(logBulk, 500)
	}
	if len(ttBulk) > 0 {
		_, _ = TokenTransferSvc.BatchCreate(ttBulk)
	}
}
