package main

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
	"runtime"
	"time"
)

func main() {
	// 一个线程获取区块，一个线程插入log数据
	runtime.GOMAXPROCS(2)
	for{
		eth.GetWaitSyncBlocks()
		time.Sleep(eth.LoopSleepDuration)
	}
}
