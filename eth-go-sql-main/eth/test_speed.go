package eth

import (
	"context"
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
	"github.com/panjf2000/ants/v2"
	"math/big"
	"os"
	"sync"
	"time"
)

func TestSpeed(url string, startBlockNumber,endBlockNumber int64){
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println("Error Client:"+ url)
		os.Exit(-1)
	}
	defer client.Close()
	defer ants.Release()

	startTime := time.Now()

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(LimitConcurrency, func(num interface{}) {
		n := num.(int64)
		getBlockDataWithoutDb(client, n)
		wg.Done()
	})
	defer p.Release()

	for i:=startBlockNumber;i<=endBlockNumber;i++{
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()

	usedTime := time.Since(startTime)
	fmt.Printf("100 Block Used Time:%v \n", usedTime)
}

func getBlockDataWithoutDb(client *ethclient.Client, number int64) {
	fmt.Println("Block: %v", number)

	blockNum := big.NewInt(number)
	block, _,_, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		fmt.Println("block number: ", blockNum, "===>", err)
		return
	}

	rpcTxCount := len(block.Transactions())
	if rpcTxCount == 0 {
		// 调用接口返回的交易数量为 0 的不作处理
		return
	}

	// 已入库交易数 不等于 调用接口返回的交易数 时， 就逐个检查tx中的Logs数量
	for _, tx := range block.Transactions() {
		_, _, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			fmt.Println(err)
		}
	}
}

