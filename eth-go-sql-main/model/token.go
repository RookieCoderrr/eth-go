package model

type Token struct {
	Address        string  `json:"address" gorm:"primaryKey;index;column:address;type:varchar(64);not null"`
	Symbol         string  `json:"symbol" gorm:"column:symbol;type:varchar(64);not null"`
	Name           string  `json:"name" gorm:"column:name;type:varchar(64);not null"`
	Decimals       int16   `json:"decimals" gorm:"column:decimals;type:int2;not null"`
	IsErc20        bool    `json:"is_erc20" gorm:"column:is_erc20;type:bool;not null"`
	IsErc721       bool    `json:"is_erc721" gorm:"column:is_erc721;type:bool;not null"`
	TotalSupply    float64 `json:"total_supply" gorm:"column:total_supply;type:numeric(64,10);not null"`
	BlockTimestamp int     `json:"block_timestamp" gorm:"column:block_timestamp;type:int4;not null;check:(\"block_timestamp\" >= 0)"`
	BlockNumber    int     `json:"block_number" gorm:"column:block_number;type:int4;not null;check:(\"block_number\" >= 0)"`
	BlockHash      string  `json:"block_hash" gorm:"column:block_hash;type:varchar(128);not null"`
}

func (m *Token) TableName() string {
	return "tokens"
}

func NewToken() *Token {
	return &Token{}
}
