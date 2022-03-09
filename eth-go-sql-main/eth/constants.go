package eth

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"os"
	"strconv"
	"time"

	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
	"git.cocosbcx.net/bigdata/eth-go-sql/service"
)

var (

	CheckBlockNumberCount, _ = strconv.ParseInt(os.Getenv("CHECK_BLOCK_NUMBER_COUNT"), 10, 64)

	loopIntervalSecond = os.Getenv("LOOP_INTERVAL_SECOND")
	sec, _             = strconv.Atoi(loopIntervalSecond)
	LoopSleepDuration  = time.Second * time.Duration(sec)

	ProviderAddr = os.Getenv("SYNC_PROVIDER_ADDR")

	TransferEventTopic              = os.Getenv("SYNC_TRANSFER_EVENT_TOPIC")
	Erc1155TransferSingleEventTopic = os.Getenv("SYNC_ERC1155_TRANSFER_SINGLE_EVENT_TOPIC")
	Erc1155TransferBatchEventTopic  = os.Getenv("SYNC_ERC1155_TRANSFER_BATCH_EVENT_TOPIC")

	LimitConcurrency, _ = strconv.Atoi(os.Getenv("SYNC_LIMIT_CONCURRENCY"))

	// 区块服务
	blockRepo = repo.NewBlockRepository(config.GlobalDB, config.GlobalLogger)
	BlockSvc  = service.NewBlockService(blockRepo)

	// 交易服务
	txRepo = repo.NewTransactionRepository(config.GlobalDB, config.GlobalLogger)
	TransactionSvc  = service.NewTransactionService(txRepo)

	// 账户服务
	blcRepo = repo.NewBalanceRepository(config.GlobalDB, config.GlobalLogger)
	BalanceSvc  = service.NewBalanceService(blcRepo)

	// 合约服务
	ctcRepo = repo.NewContractRepository(config.GlobalDB, config.GlobalLogger)
	ContractSvc  = service.NewContractService(ctcRepo)

	// 日志服务
	logRepo = repo.NewLogRepository(config.GlobalDB, config.GlobalLogger)
	LogSvc  = service.NewLogService(logRepo)

	// 转账记录服务
	ttRepo = repo.NewTokenTransferRepository(config.GlobalDB, config.GlobalLogger)
	TokenTransferSvc  = service.NewTokenTransferService(ttRepo)

	// Redis缓存服务
	RedisRepo = repo.NewRedisRepository(config.GlobalRedis)
)
