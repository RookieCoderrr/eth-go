package mq

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"sync"
	"time"
)

var PROTOCOL string
var LOGIN string
var PASSWORD string
var HOST string
var PORT string = "5671"
var VIRTUALHOST string = "/"

var conn *amqp.Connection
var channel *amqp.Channel
var chLock sync.RWMutex
var connLock sync.RWMutex

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		err = godotenv.Load("../config.env")
		if err != nil {
			logrus.Fatal("read config.env file error: ", err)
		}
	}

	PROTOCOL = os.Getenv("MQ_PROTOCOL")
	LOGIN = os.Getenv("MQ_LOGIN")
	PASSWORD = os.Getenv("MQ_PASSWORD")
	HOST = os.Getenv("MQ_HOST")
	PORT = os.Getenv("MQ_PORT")
}

func getConn(reConnection bool) *amqp.Connection {
	if !reConnection {
		connLock.RLock()
		if conn != nil && !conn.IsClosed() {
			connLock.RUnlock()
			return conn
		}
		connLock.RUnlock()
	}

	connLock.Lock()
	defer connLock.Unlock()
	if conn != nil && !conn.IsClosed() {
		_ = conn.Close()
	}
	// 创建链接
	url := fmt.Sprintf("%s://%s:%s@%s:%s%s", PROTOCOL, LOGIN, PASSWORD, HOST, PORT, VIRTUALHOST)
	var err error
	for {
		conn, err = amqp.Dial(url)
		if err != nil {
			fmt.Println("Mq Failed to open Connection,Err:", err.Error())
			// 5秒重连
			time.Sleep(5 * time.Second)
			if conn != nil {
				_ = conn.Close()
			}
			fmt.Printf("MQ 链接成功: %s://%s:%s@%s:%s%s\n", PROTOCOL, LOGIN, PASSWORD, HOST, PORT, VIRTUALHOST)
			continue
		}
		return conn
	}
}

func getChannel(reConnection bool) *amqp.Channel {
	if !reConnection {
		chLock.RLock()
		if channel != nil && !conn.IsClosed() {
			chLock.RUnlock()
			return channel
		}
		chLock.RUnlock()
	}
	// 打开一个通道
	chLock.Lock()
	defer chLock.Unlock()
	for {
		ch, err := getConn(reConnection).Channel()
		if err != nil {
			fmt.Println("Mq Failed to open Channel,Err:", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		channel = ch
		return channel
	}
}

func PublishTxLog(exchange string, logData []byte, closed bool) {
	publishData := amqp.Publishing{
		Expiration:  "10000", //10秒超时
		ContentType: "text/json",
		Body:        logData,
		Timestamp:   time.Now(),
	}

	// 指定交换机发布消息
	ch := getChannel(closed)
	err := ch.Publish(exchange, "", false, false, publishData)
	if err != nil {
		fmt.Println("Publish Error:", err.Error())
		_ = ch.Close()
		// 强制更新channel
		PublishTxLog(exchange, logData, true)
	}
}
