package main

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
	"os"
	"strconv"
	"strings"
)

func RunRepair() {
	if len(os.Args)==2{
		startBlockNumber,err := strconv.ParseInt(strings.TrimSpace(os.Args[1]),10,64)
		if err!=nil{
			config.GlobalLogger.Error(err)
			os.Exit(1)
		}
		eth.CheckDataByBlockNumber("",startBlockNumber, startBlockNumber)
	}
	if len(os.Args)==3{
		startBlockNumber,err := strconv.ParseInt(strings.TrimSpace(os.Args[1]),10,64)
		endBlockNumber,err1 := strconv.ParseInt(strings.TrimSpace(os.Args[2]),10,64)
		if err!=nil || err1!=nil{
			config.GlobalLogger.Error(err)
			os.Exit(1)
		}
		eth.CheckDataByBlockNumber("",startBlockNumber, endBlockNumber)
	}
	if len(os.Args)==4{
		clientUrl := os.Args[1]
		startBlockNumber,err := strconv.ParseInt(strings.TrimSpace(os.Args[2]),10,64)
		endBlockNumber,err1 := strconv.ParseInt(strings.TrimSpace(os.Args[3]),10,64)
		if err!=nil || err1!=nil{
			config.GlobalLogger.Error(err)
			os.Exit(1)
		}
		eth.CheckDataByBlockNumber(clientUrl, startBlockNumber, endBlockNumber)
	}
}
