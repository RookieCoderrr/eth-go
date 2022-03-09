package model

import "strings"

type Contract struct {
	Address         string `json:"address" gorm:"primaryKey;index;column:address;type:varchar(64);not null"`
	Name            string `json:"name" gorm:"column:name;type:varchar(128)"`
	Bytecode        string `json:"bytecode" gorm:"column:bytecode;type:text;not null"`
	Abi             string `json:"abi" gorm:"column:abi;type:text"`
	Runs            int    `json:"runs" gorm:"column:runs;type:int4;not null;check:(\"runs\" >= 0)"`
	BytecodeHash    string `json:"bytecode_hash" gorm:"column:bytecode_hash;type:varchar(128)"`
	Source          string `json:"source" gorm:"column:source;type:text"`
	Compiler        string `json:"compiler" gorm:"column:compiler;type:varchar(128)"`
	Library         string `json:"library" gorm:"column:library;type:varchar(1024)"`
	ConstructorArgs string `json:"constructor_args" gorm:"column:constructor_args;type:text"`
	BlockNumber     int    `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
	Timestamp       int    `json:"timestamp" gorm:"column:timestamp;type:int4;not null;check:(\"timestamp\" >= 0)"`
	Creator         string `json:"creator" gorm:"column:creator;type:varchar(64)"`
	TransactionHash string `json:"transaction_hash" gorm:"column:transaction_hash;type:varchar(128)"`
}

func (m *Contract) TableName() string {
	return "contracts"
}

func NewContract() *Contract {
	return &Contract{}
}

func NewContractFromTransactionAttrs(tx *Transaction) *Contract {
	return &Contract{
		Address:         strings.ToLower(tx.ContractAddress),
		Creator:         strings.ToLower(tx.FromAddress),
		TransactionHash: strings.ToLower(tx.TransactionHash),
		Bytecode:        tx.Input,
		BlockNumber:     tx.BlockNumber,
		Timestamp:       tx.Timestamp,
		//Name:            "",
		//Abi:             "",
		//Runs:            0,
		//BytecodeHash:    "",
		//Source:          "",
		//Compiler:        "",
		//Library:         "",
		//ConstructorArgs: "",
	}
}
