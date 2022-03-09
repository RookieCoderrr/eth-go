package model

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

type Log struct {
	ID               int    `json:"id" gorm:"primaryKey;column:id;type:int4;not null;auto_increment"`
	Address          string `json:"address" gorm:"column:address;type:varchar(64);not null"`
	BlockHash        string `json:"block_hash" gorm:"column:block_hash;type:varchar(128);not null"`
	EventName        string `json:"event_name" gorm:"column:event_name;type:varchar(128);not null"`
	BlockNumber      int    `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
	Data             string `json:"data" gorm:"column:data;type:text;not null"`
	LogIndex         int    `json:"log_index" gorm:"column:log_index;type:int4;not null;check:(\"log_index\" >= 0)"`
	Removed          bool   `json:"removed" gorm:"column:removed;type:bool;not null"`
	Topics           string `json:"topics" gorm:"column:topics;type:text;not null"`
	TransactionHash  string `json:"transaction_hash" gorm:"column:transaction_hash;type:varchar(128);not null"`
	TransactionIndex int    `json:"transaction_index" gorm:"column:transaction_index;type:int4;not null;check:(\"transaction_index\" >= 0)"`
	Args             string `json:"args" gorm:"column:args;type:jsonb"`
	Topic0           string `json:"topic_0" gorm:"column:topic0;type:varchar(128);not null"`
}

func (m *Log) TableName() string {
	return "tx_logs"
}

func NewLog() *Log {
	return &Log{}
}

func NewLogFromEthTypes(lg *types.Log) *Log {
	topic0 := ""
	if len(lg.Topics) > 0 {
		topic0 = strings.ToLower(lg.Topics[0].Hex())
	}
	topics, _ := json.Marshal(lg.Topics)
	return &Log{
		Address:          strings.ToLower(lg.Address.Hex()),
		BlockHash:        strings.ToLower(lg.BlockHash.Hex()),
		BlockNumber:      int(lg.BlockNumber),
		Data:             common.Bytes2Hex(lg.Data),
		LogIndex:         int(lg.Index),
		Removed:          lg.Removed,
		Topics:           string(topics),
		TransactionHash:  strings.ToLower(lg.TxHash.Hex()),
		TransactionIndex: int(lg.TxIndex),
		Args:             "{}",
		Topic0:           topic0,
	}
}
