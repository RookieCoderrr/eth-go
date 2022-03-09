package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"strconv"
	"strings"
)

type Transaction struct {
	TransactionHash   string  `json:"transaction_hash" gorm:"primaryKey;index;column:transaction_hash;type:varchar(128);not null"`
	BlockHash         string  `json:"block_hash" gorm:"column:block_hash;type:varchar(128);not null"`
	BlockNumber       int     `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
	TransactionIndex  uint32  `json:"transaction_index" gorm:"column:transaction_index;type:int4;not null;check:(\"transaction_index\" >= 0)"`
	TxType            string  `json:"tx_type" gorm:"column:tx_type;type:text;not null"`
	FromAddress       string  `json:"from_address" gorm:"column:from_address;type:varchar(64);not null"`
	ToAddress         string  `json:"to_address" gorm:"column:to_address;type:varchar(64)"`
	Gas               int     `json:"gas" gorm:"column:gas;type:int8;not null;check:(\"gas\" >= 0)"`
	CumulativeGasUsed int     `json:"cumulative_gas_used" gorm:"column:cumulative_gas_used;type:int4;not null;check:(\"cumulative_gas_used\" >= 0)"`
	GasUsed           int     `json:"gas_used" gorm:"column:gas_used;type:int4;not null;check:(\"gasUsed\" >= 0)"`
	GasPrice          float64 `json:"gas_price" gorm:"column:gas_price;type:numeric(64,0);not null"`
	Nonce             int     `json:"nonce" gorm:"column:nonce;type:int4;not null;check:(\"nonce\" >= 0)"`
	Value             float64 `json:"value" gorm:"column:value;type:numeric(64,0);not null"`
	R                 string  `json:"r" gorm:"column:r;type:varchar(128);not null"`
	S                 string  `json:"s" gorm:"column:s;type:varchar(128);not null"`
	V                 int     `json:"v" gorm:"column:v;type:int4;not null;check:(\"v\" >= 0)"`
	Status            int     `json:"status" gorm:"column:status;type:int2;not null"`
	Timestamp         int     `json:"timestamp" gorm:"column:timestamp;type:int4;not null;check:(\"timestamp\" >= 0)"`
	LogsBloom         string  `json:"logs_bloom" gorm:"column:logs_bloom;type:text;not null"`
	ContractAddress   string  `json:"contract_address" gorm:"column:contract_address;type:varchar(128)"`
	LogCount          int     `json:"log_count" gorm:"column:log_count;type:int4;not null"`
	Input             string  `json:"input" gorm:"column:input;type:text"`
}

func (m *Transaction) TableName() string {
	return "transcations"
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func NewTransactionFromEthTypes(block *types.Block, fromAddr string, tx *types.Transaction, receipt *types.Receipt) *Transaction {
	bloom, _ := receipt.Bloom.MarshalText()
	bloomStr := string(bloom)

	v, r, s := tx.RawSignatureValues()
	vv, _ := strconv.Atoi(v.String())

	toAddr := tx.To()
	toStr := ""
	if toAddr != nil {
		toStr = toAddr.Hex()
	}
	return &Transaction{
		TransactionHash:   strings.ToLower(tx.Hash().Hex()),
		BlockHash:         strings.ToLower(block.Hash().Hex()),
		BlockNumber:       int(block.Number().Int64()),
		TransactionIndex:  uint32(receipt.TransactionIndex),
		TxType:            strconv.Itoa(int(tx.Type())),
		FromAddress:       strings.ToLower(fromAddr),
		ToAddress:         strings.ToLower(toStr),
		Gas:               int(tx.Gas()),
		CumulativeGasUsed: int(receipt.CumulativeGasUsed),
		GasUsed:           int(receipt.GasUsed),
		GasPrice:          float64(tx.GasPrice().Int64()),
		Nonce:             int(tx.Nonce()),
		Value:             float64(tx.Value().Int64()),
		R:                 r.String(),
		S:                 s.String(),
		V:                 vv,
		Status:            int(receipt.Status),
		Timestamp:         int(block.Time()),
		LogsBloom:         bloomStr,
		ContractAddress:   strings.ToLower(receipt.ContractAddress.Hex()),
		LogCount:          len(receipt.Logs),
		Input:             "",
	}
}
