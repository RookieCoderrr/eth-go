package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type ILogRepository interface {
	Create(log *model.Log) (*model.Log, error)
	Delete(log *model.Log) (*model.Log, error)
	Update(log *model.Log) (*model.Log, error)
	Query(conditions interface{}) ([]*model.Log, error)
	Save(log *model.Log) (*model.Log, error)
	BatchCreate(bulk []*model.Log) (bool, error)
}

type LogRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *LogRepository) BatchCreate(bulk []*model.Log,batchSize int) (bool, error) {
	err := b.db.Model(model.NewLog()).CreateInBatches(bulk, batchSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewLogRepository(db *gorm.DB, logger *logrus.Logger) *LogRepository {
	return &LogRepository{db: db, logger: logger}
}

func (b *LogRepository) Create(log *model.Log) (*model.Log, error) {
	err := b.db.Create(&log).Error
	if err != nil {
		b.logger.Error("LogRepository.Create error: ", err)
		return nil, err
	}
	return log, nil
}

func (b *LogRepository) Delete(log *model.Log) (*model.Log, error) {
	panic("implement me")
}

func (b *LogRepository) Update(log *model.Log) (*model.Log, error) {
	panic("implement me")
}

func (b *LogRepository) Query(conditions interface{}) ([]*model.Log, error) {
	var result []*model.Log
	err := b.db.Where(conditions).Find(&result).Error
	if err != nil {
		b.logger.Error("LogRepository.Query error: ", err)
		return nil, err
	}
	return result, nil
}

func (b *LogRepository) Save(log *model.Log) (*model.Log, error) {
	panic("implement me")
}

func (b *LogRepository) QueryByBlockNumber(blockNumber int64)([]*model.Log, error){
	var result []*model.Log
	err := b.db.Model(model.Log{}).Where("block_number=?", blockNumber).Scan(&result).Error
	return result,err
}
