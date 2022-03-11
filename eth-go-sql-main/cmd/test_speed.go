package main

import (
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/eth"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args)==3{
		url := os.Args[1]
		startBlockNumber,err := strconv.ParseInt(strings.TrimSpace(os.Args[2]),10,64)
		if err!=nil {
			fmt.Println(err)
			os.Exit(1)
		}
		eth.TestSpeed(url, startBlockNumber, startBlockNumber+100)
	}
}
