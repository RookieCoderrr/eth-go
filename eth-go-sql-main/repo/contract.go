package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
)

type IContractRepository interface {
	Create(contract *model.Contract) (*model.Contract, error)
	Delete(contract *model.Contract) (*model.Contract, error)
	Update(contract *model.Contract) (*model.Contract, error)
	Query(conditions interface{}) ([]*model.Contract, error)
	Save(contract *model.Contract) (*model.Contract, error)
	BatchCreate(bulk []*model.Contract) (bool, error)
}

type ContractRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *ContractRepository) BatchCreate(bulk []*model.Contract) (bool, error) {
	err := b.db.Model(model.NewContract()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewContractRepository(db *gorm.DB, logger *logrus.Logger) *ContractRepository {
	return &ContractRepository{db: db, logger: logger}
}

func (b *ContractRepository) Create(contract *model.Contract) (*model.Contract, error) {
	err := b.db.Create(&contract).Error
	if err != nil {
		//b.logger.Error("ContractRepository.Create error: ", err)
		return nil, err
	}
	return contract, nil
}

func (b *ContractRepository) Delete(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}

func (b *ContractRepository) Update(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}

func (b *ContractRepository) Query(conditions interface{}) ([]*model.Contract, error) {
	panic("implement me")
}

func (b *ContractRepository) Save(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}
