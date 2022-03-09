package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type ITokenService interface {
	Create(token *model.Token) (*model.Token, error)
	Delete(token *model.Token) (*model.Token, error)
	Update(token *model.Token) (*model.Token, error)
	Query(conditions interface{}) ([]*model.Token, error)
	Save(token *model.Token) (*model.Token, error)
	BatchCreate(bulk []*model.Token) (bool, error)
}

type TokenService struct {
	tokenRepo *repo.TokenRepository
}

func NewTokenService(tokenRepo *repo.TokenRepository) *TokenService {
	return &TokenService{tokenRepo: tokenRepo}
}

func (b *TokenService) BatchCreate(bulk []*model.Token) (bool, error) {
	return b.tokenRepo.BatchCreate(bulk)
}

func (b *TokenService) Create(token *model.Token) (*model.Token, error) {
	inst, err := b.tokenRepo.Create(token)
	if err != nil {
		return nil, err
	}
	return inst, err
}

func (b *TokenService) Delete(token *model.Token) (*model.Token, error) {
	panic("implement me")
}

func (b *TokenService) Update(token *model.Token) (*model.Token, error) {
	panic("implement me")
}

func (b *TokenService) Query(conditions interface{}) ([]*model.Token, error) {
	panic("implement me")
}

func (b *TokenService) Save(token *model.Token) (*model.Token, error) {
	panic("implement me")
}
