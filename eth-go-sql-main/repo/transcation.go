package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"

	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type ITransactionRepository interface {
	Create(tx *model.Transaction) (*model.Transaction, error)
	Delete(tx *model.Transaction) (*model.Transaction, error)
	Update(tx *model.Transaction) (*model.Transaction, error)
	Query(conditions interface{}) ([]*model.Transaction, error)
	Save(tx *model.Transaction) (*model.Transaction, error)
	BatchCreate(bulk []*model.Transaction) (bool, error)
	GetTransactionByHash(value string) (*model.Transaction, error)
}

type TransactionRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *TransactionRepository) GetTransactionByHash(value string) (*model.Transaction, error) {
	var inst *model.Transaction
	err := b.db.Where("transaction_hash = ?", strings.ToLower(value)).First(&inst).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		b.logger.Error("TransactionRepository.GetTransactionByHash error: ", err)
		return nil, err
	}
	return inst, nil
}

func (b *TransactionRepository) BatchCreate(bulk []*model.Transaction) (bool, error) {
	err := b.db.Model(model.NewTransaction()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewTransactionRepository(db *gorm.DB, logger *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{db: db, logger: logger}
}

func (b *TransactionRepository) Create(tx *model.Transaction) (*model.Transaction, error) {
	err := b.db.Create(&tx).Error
	if err != nil {
		b.logger.Errorf("TransactionRepository【%d】:【%s】.Create error: %v",tx.BlockNumber,tx.TransactionHash, err)
		return nil, err
	}
	return tx, nil
}

func (b *TransactionRepository) QueryByBlockNumber(blockNumber int64) ([]*model.Transaction, error) {
	var result []*model.Transaction
	err := b.db.Model(model.Transaction{}).Select("transaction_hash").Where("block_number=?", blockNumber).Scan(&result).Error
	return result,err
}

func (b *TransactionRepository) Delete(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionRepository) Update(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionRepository) Query(conditions interface{}) ([]*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionRepository) Save(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}
