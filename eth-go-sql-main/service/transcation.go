package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type ITransactionService interface {
	Create(tx *model.Transaction) (*model.Transaction, error)
	Delete(tx *model.Transaction) (*model.Transaction, error)
	Update(tx *model.Transaction) (*model.Transaction, error)
	Query(conditions interface{}) ([]*model.Transaction, error)
	Save(tx *model.Transaction) (*model.Transaction, error)
	BatchCreate(bulk []*model.Transaction) (bool, error)
	GetTransactionByHash(value string) (*model.Transaction, error)
}

type TransactionService struct {
	txRepo *repo.TransactionRepository
}

func NewTransactionService(txRepo *repo.TransactionRepository) *TransactionService {
	return &TransactionService{txRepo: txRepo}
}

func (b *TransactionService) GetTransactionByHash(value string) (*model.Transaction, error) {
	return b.txRepo.GetTransactionByHash(value)
}

func (b *TransactionService) BatchCreate(bulk []*model.Transaction) (bool, error) {
	return b.txRepo.BatchCreate(bulk)
}

func (b *TransactionService) Create(tx *model.Transaction) (*model.Transaction, error) {
	inst, err := b.txRepo.Create(tx)
	if err != nil {
		return nil, err
	}
	return inst, err
}

func (b *TransactionService) Delete(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionService) Update(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionService) Query(conditions interface{}) ([]*model.Transaction, error) {
	panic("implement me")
}

func (b *TransactionService) Save(tx *model.Transaction) (*model.Transaction, error) {
	panic("implement me")
}
