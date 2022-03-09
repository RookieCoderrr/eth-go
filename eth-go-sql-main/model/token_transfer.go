package model

import (
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type TokenTransfer struct {
	ID              int    `json:"id" gorm:"primaryKey;column:id;type:int4;not null"`
	TokenAddress    string `json:"token_address" gorm:"column:token_address;type:varchar(64);not null"`
	FromAddress     string `json:"from_address" gorm:"column:from_address;type:varchar(64);not null"`
	ToAddress       string `json:"to_address" gorm:"column:to_address;type:varchar(64);not null"`
	Value           string `json:"value" gorm:"column:value;type:numeric(64,0);not null"`
	TransactionHash string `json:"transaction_hash" gorm:"column:transaction_hash;type:varchar(128);not null"`
	LogIndex        int    `json:"log_index" gorm:"column:log_index;type:int4;not null;check:(\"log_index\" >= 0)"`
	BlockHash       string `json:"block_hash" gorm:"column:block_hash;type:varchar(128);not null"`
	BlockNumber     int    `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
	EventName       string `json:"event_name" gorm:"column:event_name;type:varchar(64)"`
	Topic           string `json:"topic" gorm:"column:topic;type:varchar(128)"`
	Timestamp       int    `json:"timestamp" gorm:"column:timestamp;type:int4;check:(\"timestamp\" >= 0)"`
	TransferType    int    `json:"transfer_type" gorm:"column:transfer_type;type:int4;not null"`
}

func (m *TokenTransfer) TableName() string {
	return "token_transfers"
}

func NewTokenTransfer() *TokenTransfer {
	return &TokenTransfer{}
}

func NewTokenTransferFromFields(tx *Transaction, lg *types.Log, topicHash, eventName, fromAddr, toAddr string, value *big.Int) *TokenTransfer {
	fa := ""
	if len(fromAddr) > 2 {
		fa = utils.TrimLeftZeroes(fromAddr[2:])
	}
	ta := ""
	if len(toAddr) > 2 {
		ta = utils.TrimLeftZeroes(toAddr[2:])
	}
	return &TokenTransfer{
		Topic:           strings.ToLower(topicHash),
		TokenAddress:    strings.ToLower(lg.Address.Hex()),
		FromAddress:     strings.ToLower(fmt.Sprintf("0x%v", fa)),
		ToAddress:       strings.ToLower(fmt.Sprintf("0x%v", ta)),
		EventName:       eventName,
		Value:           value.String(),
		TransactionHash: strings.ToLower(tx.TransactionHash),
		LogIndex:        int(lg.Index),
		Timestamp:       tx.Timestamp,
		BlockHash:       strings.ToLower(tx.BlockHash),
		BlockNumber:     tx.BlockNumber,
	}
}
