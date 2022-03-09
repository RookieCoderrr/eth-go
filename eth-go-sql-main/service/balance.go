package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type IBalanceService interface {
	Create(balance *model.Balance) (*model.Balance, error)
	Delete(balance *model.Balance) (*model.Balance, error)
	Update(balance *model.Balance) (*model.Balance, error)
	GetQueryset(conditions interface{}) ([]*model.Balance, error)
	Save(balance *model.Balance) (*model.Balance, error)
	QueryByAddress(address string) (*model.Balance, error)
	BatchCreate(bulk []*model.Balance) (bool, error)
}

type BalanceService struct {
	balanceRepo *repo.BalanceRepository
}

func NewBalanceService(balanceRepo *repo.BalanceRepository) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo}
}

func (b *BalanceService) BatchCreate(bulk []*model.Balance) (bool, error) {
	return b.balanceRepo.BatchCreate(bulk)
}

func (b *BalanceService) QueryByAddress(address string) (*model.Balance, error) {
	return b.balanceRepo.QueryByAddress(address)
}

func (b *BalanceService) Create(balance *model.Balance) (*model.Balance, error) {
	return b.balanceRepo.Create(balance)
}

func (b *BalanceService) Delete(balance *model.Balance) (*model.Balance, error) {
	panic("implement me")
}

func (b *BalanceService) Update(balance *model.Balance) (*model.Balance, error) {
	return b.balanceRepo.Update(balance)
}

func (b *BalanceService) GetQueryset(conditions interface{}) ([]*model.Balance, error) {
	return b.balanceRepo.GetQueryset(conditions)
}

func (b *BalanceService) Save(balance *model.Balance) (*model.Balance, error) {
	return b.balanceRepo.Save(balance)
}
