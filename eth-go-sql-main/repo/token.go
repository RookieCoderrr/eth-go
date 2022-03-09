package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type ITokenRepository interface {
	Create(token *model.Token) (*model.Token, error)
	Delete(token *model.Token) (*model.Token, error)
	Update(token *model.Token) (*model.Token, error)
	Query(conditions interface{}) ([]*model.Token, error)
	Save(token *model.Token) (*model.Token, error)
	BatchCreate(bulk []*model.Token) (bool, error)
}

type TokenRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *TokenRepository) BatchCreate(bulk []*model.Token) (bool, error) {
	err := b.db.Model(model.NewToken()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewTokenRepository(db *gorm.DB, logger *logrus.Logger) *TokenRepository {
	return &TokenRepository{db: db, logger: logger}
}

func (b *TokenRepository) Create(token *model.Token) (*model.Token, error) {
	err := b.db.Create(&token).Error
	if err != nil {
		b.logger.Error("TokenRepository.Create error: ", err)
		return nil, err
	}
	return token, nil
}

func (b *TokenRepository) Delete(token *model.Token) (*model.Token, error) {
	panic("implement me")
}

func (b *TokenRepository) Update(token *model.Token) (*model.Token, error) {
	panic("implement me")
}

func (b *TokenRepository) Query(conditions interface{}) ([]*model.Token, error) {
	panic("implement me")
}

func (b *TokenRepository) Save(token *model.Token) (*model.Token, error) {
	panic("implement me")
}
