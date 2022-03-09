package main

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/cmd"
	"git.cocosbcx.net/bigdata/eth-go-sql/config"
)

func main () {
	config.Init()
	cmd.RunProducer()
}
