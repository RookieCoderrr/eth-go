package model

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/pkg/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

type Block struct {
	Number           *int    `json:"number" gorm:"primaryKey;column:number;type:int4;not null;check:(\"number\" >= 0)"`
	Hash             string  `json:"hash" gorm:"index:blocks_hash_b84f947b;column:hash;type:varchar(128);not null"`
	Difficulty       float64 `json:"difficulty" gorm:"column:difficulty;type:numeric(64,0);not null"`
	ExtraData        string  `json:"extra_data" gorm:"column:extra_data;type:text;not null"`
	GasLimit         int     `json:"gas_limit" gorm:"column:gas_limit;type:int4;not null;check:(\"gas_limit\" >= 0)"`
	GasUsed          int     `json:"gas_used" gorm:"column:gas_used;type:int4;not null;check:(\"gas_used\" >= 0)"`
	LogsBloom        string  `json:"logs_bloom" gorm:"column:logs_bloom;type:text;not null"`
	Miner            string  `json:"miner" gorm:"index:blocks_miner_4c97a59a;column:miner;type:varchar(64);not null"`
	MixHash          string  `json:"mix_hash" gorm:"index:blocks_mix_hash_f2ebab47;column:mix_hash;type:varchar(128);not null"`
	Nonce            string  `json:"nonce" gorm:"column:nonce;type:varchar(128);not null"`
	ParentHash       string  `json:"parent_hash" gorm:"column:parent_hash;type:varchar(128);not null"`
	ReceiptsRoot     string  `json:"receipts_root" gorm:"column:receipts_root;type:varchar(128);not null"`
	Sha3Uncles       string  `json:"sha3_uncles" gorm:"column:sha3_uncles;type:varchar(128);not null"`
	Size             int     `json:"size" gorm:"column:size;type:int4;not null;check:(\"size\" >= 0)"`
	StateRoot        string  `json:"state_root" gorm:"column:state_root;type:varchar(128);not null"`
	Timestamp        int     `json:"timestamp" gorm:"column:timestamp;type:int4;not null;check:(\"timestamp\" >= 0)"`
	TotalDifficulty  string  `json:"total_difficulty" gorm:"column:total_difficulty;type:numeric(64,0);not null"`
	TransactionsRoot string  `json:"transactions_root" gorm:"column:transactions_root;type:varchar(128);not null"`
	Uncles           string  `json:"uncles" gorm:"column:uncles;type:text;not null"`
	TransactionCount int     `json:"transaction_count" gorm:"column:transaction_count;type:int4;not null;check:(\"transaction_count\" >= 0)"`
}

func (m *Block) TableName() string {
	return "blocks"
}

func NewBlock() *Block {
	return &Block{}
}

func NewBlockFromEthTypes(eb *types.Block, totalDifficulty *hexutil.Big) *Block {
	tdValue := utils.HexToBigInt(totalDifficulty.String()).String()
	header := eb.Header()
	bloom, _ := header.Bloom.MarshalText()
	bloomStr := string(bloom)

	num := int(header.Number.Int64())
	return &Block{
		Number:           &num,
		Hash:             strings.ToLower(header.Hash().Hex()),
		Difficulty:       float64(header.Difficulty.Int64()),
		ExtraData:        hexutil.Bytes(header.Extra).String(),
		GasLimit:         int(header.GasLimit),
		GasUsed:          int(header.GasUsed),
		LogsBloom:        bloomStr,
		Miner:            header.Coinbase.String(),
		MixHash:          strings.ToLower(header.MixDigest.Hex()),
		Nonce:            hexutil.EncodeUint64(header.Nonce.Uint64()),
		ParentHash:       strings.ToLower(eb.ParentHash().Hex()),
		ReceiptsRoot:     eb.ReceiptHash().Hex(),
		Sha3Uncles:       strings.ToLower(header.UncleHash.Hex()),
		Size:             int(eb.Size()),
		StateRoot:        header.Root.Hex(),
		Timestamp:        int(eb.Time()),
		TotalDifficulty:  tdValue,
		TransactionsRoot: strings.ToLower(header.TxHash.Hex()),
		Uncles:           "",
		TransactionCount: len(eb.Transactions()),
	}
}
