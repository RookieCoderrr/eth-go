package main

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
)

func main() {
	// 单独的线程来处理失败的交易数据
	go eth.CheckRedisFailedTransaction()
	// 单独的线程处理补块程序
	//for{
	//	eth.CheckEthBlockData()
	//	time.Sleep(time.Minute)
	//}
}
