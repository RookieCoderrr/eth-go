package repo

import (
	"fmt"
	"git.cocosbcx.net/bigdata/eth-go-sql/config"
	"git.cocosbcx.net/bigdata/eth-go-sql/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IBlockRepository interface {
	Create(inst *model.Block) (*model.Block, error)
	Delete(number int) (bool, error)
	Update(inst *model.Block) (*model.Block, error)
	GetQueryset(conditions interface{}, orderBy interface{}) ([]*model.Block, error)
	Save(inst *model.Block) (*model.Block, error)
	QueryByNumber(number int) (*model.Block, error)
	BatchCreate(bulk []*model.Block) (bool, error)
	GetMaxNumber() (*int, error)
}

type BlockRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (b *BlockRepository) BatchCreate(bulk []*model.Block) (bool, error) {
	err := b.db.Model(model.NewBlock()).CreateInBatches(bulk, config.BulkSize).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewBlockRepository(db *gorm.DB, logger *logrus.Logger) *BlockRepository {
	return &BlockRepository{db: db, logger: logger}
}

func (b *BlockRepository) GetMaxNumber() (*int, error) {
	var max int
	sql := fmt.Sprintf("select max(number) from %v", model.NewBlock().TableName())
	err := b.db.Raw(sql).Scan(&max).Error
	if err != nil {
		b.logger.Error("BlockRepository GetMaxNumber error: ", err,", Use BlockNumber:0")
		return &max, nil
	}
	return &max, nil
}


func (b *BlockRepository) Create(block *model.Block) (*model.Block, error) {
	err := b.db.Create(&block).Error
	if err != nil {
		b.logger.Error("BlockRepository.Create error: ", err)
		return nil, err
	}
	return block, nil
}

func (b *BlockRepository) Delete(number int) (bool, error) {
	err := b.db.Where("number = ?", number).Delete(model.NewBlock()).Error
	if err != nil {
		b.logger.Error("BlockRepository.Delete error: ", err)
		return false, err
	}
	return true, nil
}

func (b *BlockRepository) Update(inst *model.Block) (*model.Block, error) {
	return b.Save(inst)
}

func (b *BlockRepository) GetQueryset(conditions interface{}, orderBy interface{}) ([]*model.Block, error) {
	var blocks []*model.Block
	qs := b.db.Where(conditions).Find(&blocks)
	if orderBy != nil {
		qs = qs.Order(orderBy)
	}
	err := qs.Error
	if err != nil {
		b.logger.Error("BlockRepository.Query error: ", err)
		return nil, err
	}
	return blocks, nil
}

func (b *BlockRepository) Save(inst *model.Block) (*model.Block, error) {
	err := b.db.Save(&inst).Error
	if err != nil {
		b.logger.Error("BlockRepository.Save error: ", err)
		return nil, err
	}
	return inst, nil
}

func (b *BlockRepository) QueryByNumber(number int64) (*model.Block, error) {
	var inst *model.Block
	err := b.db.Where("number = ?", number).First(&inst).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		b.logger.Error("BlockRepository.QueryByNumber error: ", err)
		return nil, err
	}
	return inst, nil
}

func (b *BlockRepository) NumberExist(number int64) bool{
	var count int64
	err := b.db.Model(model.Block{}).Where("number=?", number).Count(&count).Error
	if count==0 || err!=nil{
		return false
	}
	return true
}

func (b *BlockRepository) QueryAllBlockNumber() ([]model.Block,error){
	var inst []model.Block
	err:=b.db.Model(model.Block{}).Select("number").Order("number asc").Scan(&inst).Error
	if err!=nil{
		b.logger.Error("GetAllBlockNumber Error: ", err)
		return nil, err
	}
	return inst,nil
}
