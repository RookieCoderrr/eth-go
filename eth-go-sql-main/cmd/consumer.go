package cmd

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
	"os"
	"runtime"
	"strings"
	"time"
)

func RunConsumer() {
	runtime.GOMAXPROCS(eth.LimitConcurrency)
	var specClientUrl string
	isInsertPsql := true
	if len(os.Args)>1{
		if os.Args[1]=="skipdb"{
			isInsertPsql = false
		}else{
			specClientUrl = strings.TrimSpace(os.Args[1])
		}
	}
	for {
		eth.SyncEth(specClientUrl, isInsertPsql)
		time.Sleep(eth.LoopSleepDuration)
	}
}
