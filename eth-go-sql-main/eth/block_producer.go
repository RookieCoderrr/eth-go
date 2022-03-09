package eth

import (
	"context"
	. "git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
)

// 上次记录的最后的区块号
var LastBlockNumber int64

//GetWaitSyncBlocks 顺序获取block编号放到队列里
func GetWaitSyncBlocks() {
	defer func() {
		if err := recover(); err != nil {
			GlobalLogger.Errorf("M%v\n", err)
		}
	}()

	queueLength := RedisRepo.Llen(GlobalBlockSeqQueue)
	if queueLength > 0 {
		GlobalLogger.Infof("%s Length:%d", GlobalBlockSeqQueue, queueLength)
		if queueLength > 10 {
			return
		}
	}

	// 获取最快的服务节点
	node := GetFasterNode()
	client, err := ethclient.Dial(node.Url)
	if err != nil {
		GlobalLogger.Error(err)
		return
	}
	defer client.Close()

	// 显示当前节点同步信息
	process, err := client.SyncProgress(context.Background())
	if err != nil {
		GlobalLogger.Error(err)
		return
	}
	if process != nil {
		GlobalLogger.Infof("CurrentBlock:%d, HighestBlock:%d, KnownStates:%d, PulledStates:%d", process.CurrentBlock, process.HighestBlock, process.KnownStates, process.PulledStates)
	}
	// 当前链上的最大节点  todo: currentBlockNum = node.BlockNumber
	currentBlockNum, err := GetHeaderMaxNumber(client)
	if err != nil {
		return
	}
	var startBlockNumber int64
	if LastBlockNumber == 0 {
		// 数据库当前的节点
		startBlockNumber = GetStartBlockNumber()
	} else {
		startBlockNumber = LastBlockNumber + 1
	}

	for blockNum := startBlockNumber; blockNum <= currentBlockNum; blockNum++ {
		length, err := RedisRepo.LPush(GlobalBlockSeqQueue, blockNum)
		if err != nil {
			GlobalLogger.Fatal("Redis Error:", err)
		}
		if length%100 == 0 {
			GlobalLogger.Infof("%s Length:%d", GlobalBlockSeqQueue, length)
		}
		LastBlockNumber = blockNum
	}
}
