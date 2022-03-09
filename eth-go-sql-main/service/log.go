package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type ILogService interface {
	Create(log *model.Log) (*model.Log, error)
	Delete(log *model.Log) (*model.Log, error)
	Update(log *model.Log) (*model.Log, error)
	Query(conditions interface{}) ([]*model.Log, error)
	Save(log *model.Log) (*model.Log, error)
	BatchCreate(bulk []*model.Log) (bool, error)
}

type LogService struct {
	logRepo *repo.LogRepository
}

func NewLogService(logRepo *repo.LogRepository) *LogService {
	return &LogService{logRepo: logRepo}
}

func (b *LogService) BatchCreate(bulk []*model.Log, batchSize int) (bool, error) {
	return b.logRepo.BatchCreate(bulk,batchSize)
}

func (b *LogService) Create(log *model.Log) (*model.Log, error) {
	inst, err := b.logRepo.Create(log)
	if err != nil {
		return nil, err
	}
	return inst, err
}

func (b *LogService) Delete(log *model.Log) (*model.Log, error) {
	panic("implement me")
}

func (b *LogService) Update(log *model.Log) (*model.Log, error) {
	panic("implement me")
}

func (b *LogService) Query(conditions interface{}) ([]*model.Log, error) {
	return b.logRepo.Query(conditions)
}

func (b *LogService) Save(log *model.Log) (*model.Log, error) {
	panic("implement me")
}
