package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type ITokenTransferService interface {
	Create(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error)
	Delete(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error)
	Update(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error)
	Query(conditions interface{}) ([]*model.TokenTransfer, error)
	Save(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error)
	BatchCreate(bulk []*model.TokenTransfer) (bool, error)
}

type TokenTransferService struct {
	tokenTransferRepo *repo.TokenTransferRepository
}

func NewTokenTransferService(tokenTransferRepo *repo.TokenTransferRepository) *TokenTransferService {
	return &TokenTransferService{tokenTransferRepo: tokenTransferRepo}
}

func (b *TokenTransferService) BatchCreate(bulk []*model.TokenTransfer) (bool, error) {
	return b.tokenTransferRepo.BatchCreate(bulk)
}

func (b *TokenTransferService) Create(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error) {
	return b.tokenTransferRepo.Create(tokenTransfer)
}

func (b *TokenTransferService) Delete(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferService) Update(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferService) Query(conditions interface{}) ([]*model.TokenTransfer, error) {
	panic("implement me")
}

func (b *TokenTransferService) Save(tokenTransfer *model.TokenTransfer) (*model.TokenTransfer, error) {
	panic("implement me")
}
