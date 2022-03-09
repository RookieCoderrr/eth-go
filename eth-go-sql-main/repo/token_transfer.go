package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type ITokenTransferRepository interface {
	Create(tt *model.TokenTransfer) (*model.TokenTransfer, error)
	Delete(tt *model.TokenTransfer) (*model.TokenTransfer, error)
	Update(tt *model.TokenTransfer) (*model.TokenTransfer, error)
	Query(conditions interface{}) ([]*model.TokenTransfer, error)
	Save(tt *model.TokenTransfer) (*model.TokenTransfer, error)
	BatchCreate(bulk []*model.TokenTransfer) (bool, error)
}

type TokenTransferRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *TokenTransferRepository) BatchCreate(bulk []*model.TokenTransfer) (bool, error) {
	err := b.db.Model(model.NewTokenTransfer()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewTokenTransferRepository(db *gorm.DB, logger *logrus.Logger) *TokenTransferRepository {
	return &TokenTransferRepository{db: db, logger: logger}
}

func (b *TokenTransferRepository) Create(tt *model.TokenTransfer) (*model.TokenTransfer, error) {
	err := b.db.Create(&tt).Error
	if err != nil {
		b.logger.Error("TokenTransferRepository.Create error: ", err)
		return nil, err
	}
	return tt, nil
}

func (b *TokenTransferRepository) Delete(tt *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferRepository) Update(tt *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferRepository) Query(conditions interface{}) ([]*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferRepository) Save(tt *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}
