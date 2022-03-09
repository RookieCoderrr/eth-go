package eth

import (
	"encoding/json"
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/mq"
	"github.com/panjf2000/ants/v2"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	. "git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
)

type NodeBlockNumber struct {
	Url         string
	BlockNumber int64
}

//GetNodeHeaderBlockNumber 获取每个节点当前的区块高度
func GetNodeHeaderBlockNumber() []NodeBlockNumber {
	var result []NodeBlockNumber
	nodes := strings.Split(ProviderAddr, "|")
	for _, n := range nodes {
		client, err := ethclient.Dial(n)
		if err != nil {
			GlobalLogger.Error("Nodes client dial error: ", err)
			continue
		}
		num, err := GetHeaderMaxNumber(client)
		if err != nil {
			continue
		}
		result = append(result, NodeBlockNumber{n, num})
		client.Close()
	}
	return result
}

func GetNodeHeaderBlockNumberByUrl(url string) (*NodeBlockNumber, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		GlobalLogger.Error("Nodes client dial error: ", err)
		return nil, err
	}
	num, err := GetHeaderMaxNumber(client)
	return &NodeBlockNumber{url, num}, err
}

func GetFasterNode() *NodeBlockNumber {
	nodeList := GetNodeHeaderBlockNumber()
	var maxBlockNumber int64
	var fasterNode NodeBlockNumber
	for _, n := range nodeList {
		GlobalLogger.Infof("Node:%s, BlockNumber:%d", n.Url, n.BlockNumber)
		if n.BlockNumber > maxBlockNumber {
			maxBlockNumber = n.BlockNumber
			fasterNode = n
		}
	}
	GlobalLogger.Infof("Select FasterNode:%s, BlockNumber:%d", fasterNode.Url, fasterNode.BlockNumber)
	return &fasterNode
}

//SyncEth 同步ETH信息, endNumber为0时取线上最近的Block Number
func SyncEth(specClientUrl string, insert bool) {
	defer func() {
		if err := recover(); err != nil {
			GlobalLogger.Errorf("M%v\n", err)
		}
	}()

	var node *NodeBlockNumber
	var err error
	if specClientUrl == "" {
		node = GetFasterNode()
	} else {
		node, err = GetNodeHeaderBlockNumberByUrl(specClientUrl)
		if err != nil {
			GlobalLogger.Error("ClientUrl is invalid")
			os.Exit(1)
		}
	}

	GlobalLogger.Infof("Node:%s ,Head BlockNumber:%d", node.Url, node.BlockNumber)

	client, err := ethclient.Dial(node.Url)
	if err != nil {
		GlobalLogger.Error(err)
		return
	}
	defer client.Close()

	// 多线程处理任务
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(LimitConcurrency, func(num interface{}) {
		GetBlock(client, num.(int64), insert)
		wg.Done()
	})
	defer p.Release()
	defer ants.Release()
	for {
		// 阻塞获取队列消息，等待1秒
		values, err := RedisRepo.BRPop(time.Second, GlobalBlockSeqQueue)
		if err != nil {
			break
		}
		if values != nil && len(values) == 2 {
			blockNumber, _ := strconv.ParseInt(values[1], 10, 64)
			// 如果区块节点小于当前要抓的节点，直接跳出循环,并将区块号放回去
			if node.BlockNumber < blockNumber {
				_, _ = RedisRepo.RPush(GlobalBlockSeqQueue, blockNumber)
				break
			}
			// wg.Add 必须放在break之后，不然会无限等待
			wg.Add(1)
			_ = p.Invoke(blockNumber)
		} else {
			break
		}
	}
	wg.Wait()
}

//BatchInsertTxLogs 从redis队列中获取TxLog数据并批量插入到数据库中
func BatchInsertTxLogs(insert bool) {
	for {
		// 阻塞获取队列消息，等待1秒
		values, err := RedisRepo.BRPop(0, GlobalTxLogQueue)
		if err != nil {
			fmt.Printf("获取%s数据异常，Error:%v\n", GlobalTxLogQueue, err)
			continue
		}
		if values != nil && len(values) == 2 {
			var logBulk []*model.Log
			byteLogData := []byte(values[1])
			start := time.Now()
			mq.PublishTxLog(MqExchangeName, byteLogData, false)
			slice := time.Since(start)
			fmt.Println("区块：", logBulk[0].BlockNumber, "发送Mq用时：", slice)
			_ = json.Unmarshal(byteLogData, &logBulk)
			if insert {
				success, err := LogSvc.BatchCreate(logBulk, BulkSize)
				if !success {
					fmt.Println("插入数据未成功")
				}
				if err != nil {
					fmt.Println("插入数据报错：", err.Error())
				}
			} else {
				fmt.Println("Lost Block:", logBulk[0].BlockNumber)
			}
			slice = time.Since(start)
			fmt.Println("区块：", logBulk[0].BlockNumber, ",Logs数量：", len(logBulk), "批量500,总共用时：", slice)
		} else {
			continue
		}
	}
}

//CheckEthBlockData 自动补块程序，补最近的5000个块
func CheckEthBlockData() {
	node := GetFasterNode()
	client, err := ethclient.Dial(node.Url)
	if err != nil {
		GlobalLogger.Fatal(err)
	}
	defer client.Close()

	defer ants.Release()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(LimitConcurrency, func(num interface{}) {
		n := num.(int64)
		CheckTxAndLogData(client, n)
		wg.Done()
	})
	defer p.Release()

	endBlockNumber := node.BlockNumber
	startBlockNumber := endBlockNumber - CheckBlockNumberCount

	// 最好按照正序进行补块，这样才不会跟新建块造成冲突
	for blockNum := startBlockNumber; blockNum <= endBlockNumber; blockNum++ {
		wg.Add(1)
		_ = p.Invoke(blockNum)
	}
	wg.Wait()
}

// CheckDataByBlockNumber 手动补块程序，指定范围或其中一个
func CheckDataByBlockNumber(clientUrl string, startBlockNumber, endBlockNumber int64) {
	var client *ethclient.Client
	var err error
	if clientUrl != "" {
		client, err = ethclient.Dial(clientUrl)
	} else {
		node := GetFasterNode()
		client, err = ethclient.Dial(node.Url)
	}
	if err != nil {
		GlobalLogger.Fatal(err)
	}

	defer client.Close()

	defer ants.Release()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(LimitConcurrency, func(num interface{}) {
		n := num.(int64)
		CheckTxAndLogData(client, n)
		wg.Done()
	})
	defer p.Release()

	for i := startBlockNumber; i <= endBlockNumber; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
}

func CheckRedisFailedTransaction() {
	for {
		node := GetFasterNode()
		client, err := ethclient.Dial(node.Url)
		if err != nil {
			GlobalLogger.Error("CheckRedisFailedTransaction client dial error: ", err)
			time.Sleep(time.Second * 30)
			continue
		}
		// 5000次重连一次
		for i := 0; i < 5000; i++ {
			// redis 返回的 values 数据结构如下：
			// ["tx_failed_queue", "5500004!-_-!0xaf228a13d63c6850e1dbc3aceb9d59dfc8882ea84670ab0f1a6cbe390c1e0177"]
			values, _ := RedisRepo.BRPop(0, GlobalTxFailedQueue)
			if values != nil && len(values) == 2 {
				HandlerSingleTransactionForChecking(client, values[1])
			}
			fmt.Println("handled 5000 txs")
		}
		client.Close()
	}
}
