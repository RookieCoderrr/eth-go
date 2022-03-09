package service

import (
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"git.cocosbcx.net/bigdata/eth-go-sql/repo"
)

type IBlockService interface {
	Create(block *model.Block) (*model.Block, error)
	Delete(block *model.Block) (*model.Block, error)
	Update(block *model.Block) (*model.Block, error)
	GetQueryset(conditions interface{}, orderBy interface{}) ([]*model.Block, error)
	Save(block *model.Block) (*model.Block, error)
	QueryByNumber(number int) (*model.Block, error)
	BatchCreate(bulk []*model.Block) (bool, error)
	GetMaxNumber() (*int, error)
}

type BlockService struct {
	blockRepo *repo.BlockRepository
}

func NewBlockService(blockRepo *repo.BlockRepository) *BlockService {
	return &BlockService{blockRepo: blockRepo}
}

func (b *BlockService) GetMaxNumber() (*int, error) {
	return b.blockRepo.GetMaxNumber()
}

func (b *BlockService) BatchCreate(bulk []*model.Block) (bool, error) {
	return b.blockRepo.BatchCreate(bulk)
}

func (b *BlockService) QueryByNumber(number int64) (*model.Block, error) {
	return b.blockRepo.QueryByNumber(number)
}

func (b *BlockService) NumberExist(number int64) bool {
	return b.blockRepo.NumberExist(number)
}

func (b *BlockService) Create(block *model.Block) (*model.Block, error) {
	return b.blockRepo.Create(block)
}

func (b *BlockService) Delete(block *model.Block) (*model.Block, error) {
	panic("implement me")
}

func (b *BlockService) Update(block *model.Block) (*model.Block, error) {
	return b.blockRepo.Update(block)
}

func (b *BlockService) GetQueryset(conditions interface{}, orderBy interface{}) ([]*model.Block, error) {
	return b.blockRepo.GetQueryset(conditions, orderBy)
}

func (b *BlockService) Save(block *model.Block) (*model.Block, error) {
	return b.blockRepo.Save(block)
}

func (b *BlockService) QueryAllBlockNumber() []int{
	var result []int
	list,err :=  b.blockRepo.QueryAllBlockNumber()
	if err!=nil{
		return result
	}
	for _,item := range list{
		result = append(result, *item.Number)
	}
	return result
}
