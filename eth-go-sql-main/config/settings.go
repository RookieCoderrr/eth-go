package config

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"strconv"
	"strings"
)

var Env string
var GlobalLogger *logrus.Logger
var GlobalDB *gorm.DB
var BulkSize = 500
var GlobalRedis *redis.Client
var ChainName string
var GlobalTxFailedQueue string
var GlobalBlockSeqQueue string
var GlobalTxLogQueue string
var MqExchangeName string

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		err = godotenv.Load("../config.env")
		if err != nil {
			logrus.Fatal("read config.env file error: ", err)
		}
	}

	Env = os.Getenv("ENV")

	ChainName = os.Getenv("CHAIN_NAME")
	GlobalTxFailedQueue = ChainName+"_tx_failed_queue"
	GlobalBlockSeqQueue = ChainName+"_block_seq_queue"
	GlobalTxLogQueue = ChainName+"_tx_log_seq_queue"
	MqExchangeName =  fmt.Sprintf("%s-exchange", strings.ToLower(ChainName))

	initLogger()
	initDBConn()
	initRedisConn()
}

func initLogger() {
	writerStd := os.Stdout
	writerFile, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Fatalf("create file log.txt failed: %v", err)
	}

	GlobalLogger = logrus.New()
	GlobalLogger.SetLevel(logrus.InfoLevel)
	GlobalLogger.SetFormatter(&logrus.TextFormatter{})
	GlobalLogger.SetOutput(io.MultiWriter(writerStd, writerFile))
}

func initDBConn() {
	// db setting
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	GlobalLogger.Infoln("database info: ", dsn)
	dryRun := os.Getenv("DB_DRY_RUN") == "true"
	gormCfg := &gorm.Config{
		DryRun: dryRun,
		//SkipDefaultTransaction:true,			// 禁用事务提高插入效率
		Logger: logger.Default.LogMode(logger.Error),
	}
	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), gormCfg,
	)
	if err != nil {
		panic("database connection init error >>> " + err.Error())
	}

	maxIdleConn, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_IDLE_CONN"))
	if err != nil {
		panic("DB_POOL_MAX_IDLE_CONN value error >>> " + err.Error())
	}
	maxOpenConn, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_OPEN_CONN"))
	if err != nil {
		panic("DB_POOL_MAX_OPEN_CONN value error >>> " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("database sql db instance init error >>> " + err.Error())
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	// sqlDB.SetConnMaxLifetime(time.Minute * 30)

	if Env=="dev" && GlobalLogger.GetLevel() == logrus.DebugLevel {
		GlobalDB = db.Debug()
	} else {
		GlobalDB = db
	}
}

func initRedisConn() {
	addr := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	_ = os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")
	_, _ = strconv.Atoi(dbStr)
	GlobalRedis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", addr, port),
	})
}