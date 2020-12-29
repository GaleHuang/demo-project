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

func (t ProductTable) QueryByPriceOrNameOrDes(query string, priceLow float32, priceHigh float32) (*[]ProductTable, error) {
	result := make([]ProductTable, 0)
	query = "%" + query + "%"
	sqlRunner := services.GetMainDB()
	if query != "" {
		sqlRunner.Where("name Like ? or des Like ?", query, query)
	}
	if !(priceLow == 0 && priceHigh == 0) {
		sqlRunner.Where("price Between ? And ?", priceLow, priceHigh)
	}
	res := sqlRunner.Find(&result)
	return &result, res.Error
}
