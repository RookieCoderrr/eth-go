package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type IContractService interface {
	Create(contract *model.Contract) (*model.Contract, error)
	Delete(contract *model.Contract) (*model.Contract, error)
	Update(contract *model.Contract) (*model.Contract, error)
	Query(conditions interface{}) ([]*model.Contract, error)
	Save(contract *model.Contract) (*model.Contract, error)
	BatchCreate(bulk []*model.Contract) (bool, error)
}

type ContractService struct {
	contractRepo *repo.ContractRepository
}

func NewContractService(contractRepo *repo.ContractRepository) *ContractService {
	return &ContractService{contractRepo: contractRepo}
}

func (b *ContractService) BatchCreate(bulk []*model.Contract) (bool, error) {
	return b.contractRepo.BatchCreate(bulk)
}

func (b *ContractService) Create(contract *model.Contract) (*model.Contract, error) {
	return b.contractRepo.Create(contract)
}

func (b *ContractService) Delete(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}

func (b *ContractService) Update(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}

func (b *ContractService) Query(conditions interface{}) ([]*model.Contract, error) {
	panic("implement me")
}

func (b *ContractService) Save(contract *model.Contract) (*model.Contract, error) {
	panic("implement me")
}
