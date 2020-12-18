package table

import (
	"github.com/galehuang/demo-project/services"
)

type ProductTable struct {
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

func (t ProductTable) UpdateOne() error {
	return services.GetMainDB().Updates(&t).Error
}

func (t ProductTable) DeleteOne(id uint64) error {
	return services.GetMainDB().Where("id = ?", id).Delete(&t).Error
}

func (t ProductTable) InsertMany(tables []ProductTable) error {
	return services.GetMainDB().Create(&tables).Error
}

func (t ProductTable) QueryByNameOrDes(query string) (*[]ProductTable, error) {
	query = "%" + query + "%"
	result := make([]ProductTable, 0)
	res := services.GetMainDB().Where("name Like ? Or des Like ?", query, query).Find(&result)
	return &result, res.Error
}


func (t ProductTable) QueryByPriceOrNameOrDes(query string, priceLow float32, priceHigh float32) (*[]ProductTable, error)  {
	query = "%" + query + "%"
	result := make([]ProductTable, 0)
	res := services.GetMainDB().Where("price Between ? And ?", priceLow, priceHigh).Where("name Like ? Or des Like ?", query, query).Find(&result)
	return &result, res.Error
}
