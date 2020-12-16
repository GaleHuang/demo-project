package table

import (
	"github.com/galehuang/demo-project/services"
	"gorm.io/gorm"
)

type ProductTable struct {
	gorm.Model
	Id    uint64
	Name  string
	Des   string
	Price float64
}

func (t ProductTable) TableName() string {
	return "tb_product"
}

func (t ProductTable) QueryOneById(id int64) (*ProductTable, error) {
	result := services.GetMainDB().First(&t, id)
	return &t, result.Error
}

func (t ProductTable) InsertOne() (uint64, error) {
	result := services.GetMainDB().Create(&t)
	return t.Id, result.Error
}

func (t ProductTable) UpdateOne() error  {
	return services.GetMainDB().Updates(&t).Error
}

func (t ProductTable) DeleteOne(id uint64) error  {
	return services.GetMainDB().Where("id = ?", id).Delete(&t).Error
}

func (t ProductTable) InsertMany(tables []ProductTable) error{
	return services.GetMainDB().Create(tables).Error
}

func (t ProductTable) SearchByNameOrDes(query string) (*[]ProductTable, error)  {
	result := make([]ProductTable, 0)
	res := services.GetMainDB().Where("name Like ? or des Like ?", query, query).Find(&result)
	return &result, res.Error
}