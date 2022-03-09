package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type IBalanceRepository interface {
	Create(balance *model.Balance) (*model.Balance, error)
	Delete(balance *model.Balance) (*model.Balance, error)
	Update(balance *model.Balance) (*model.Balance, error)
	GetQueryset(conditions interface{}) ([]*model.Balance, error)
	Save(balance *model.Balance) (*model.Balance, error)
	QueryByAddress(address string) (*model.Balance, error)
	BatchCreate(bulk []*model.Balance) (bool, error)
}

type BalanceRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *BalanceRepository) BatchCreate(bulk []*model.Balance) (bool, error) {
	err := b.db.Model(model.NewBalance()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewBalanceRepository(db *gorm.DB, logger *logrus.Logger) *BalanceRepository {
	return &BalanceRepository{db: db, logger: logger}
}

func (b *BalanceRepository) QueryByAddress(address string) (*model.Balance, error) {
	var items []*model.Balance
	err := b.db.Where("address = ?", address).Find(&items).Error
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		return items[0], nil
	}
	return nil, nil
}

func (b *BalanceRepository) Create(balance *model.Balance) (*model.Balance, error) {
	err := b.db.Create(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *BalanceRepository) Delete(balance *model.Balance) (*model.Balance, error) {
	panic("implement me")
}

func (b *BalanceRepository) Update(balance *model.Balance) (*model.Balance, error) {
	return b.Save(balance)
}

func (b *BalanceRepository) GetQueryset(conditions interface{}) ([]*model.Balance, error) {
	var result []*model.Balance
	qs := b.db
	if conditions != nil {
		qs = qs.Where(conditions)
	}
	err := qs.Find(&result).Error
	if err != nil {
		b.logger.Error("BalanceRepository.Query error: ", err)
		return nil, err
	}
	return result, nil
}

func (b *BalanceRepository) Save(balance *model.Balance) (*model.Balance, error) {
	err := b.db.Save(balance).Error
	if err != nil {
		b.logger.Error("BalanceRepository.Save error: ", err)
		return nil, err
	}
	return balance, nil
}
