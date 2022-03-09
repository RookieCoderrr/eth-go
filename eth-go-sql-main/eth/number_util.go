package eth

import (
	"context"

	. "git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/ethclient"
)

// 获取要处理的区块号段
func GetStartBlockNumber() int64 {
	st, err := BlockSvc.GetMaxNumber()
	if err != nil {
		return 0
	}

	// 从库里查的最大数,加1才是开始的区块号
	start := int64(*st) + 1
	return start
}

// 取节点服务当前最大的块号
func GetHeaderMaxNumber(client *ethclient.Client) (int64,error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		GlobalLogger.Error("GetHeaderMaxNumber error: ", err)
		return 0, err
	}
	return header.Number.Int64(),nil
}
