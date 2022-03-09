package model

import "strings"

type Balance struct {
	Address     string  `json:"address" gorm:"primaryKey;index;column:address;type:varchar(64);not null"`
	EthBalance  float64 `json:"eth_balance" gorm:"column:eth_balance;type:numeric(64,0);not null"`
	BlockNumber int64   `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
}

func (m *Balance) TableName() string {
	return "balances"
}

func NewBalance() *Balance {
	return &Balance{}
}

func NewBalanceFromFields(address string, blockNumber int64, balanceValue float64) *Balance {
	return &Balance{
		Address:     strings.ToLower(address),
		BlockNumber: blockNumber,
		EthBalance:  balanceValue,
	}
}
