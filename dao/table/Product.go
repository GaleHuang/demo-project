package table

import (
	"github.com/galehuang/demo-project/services"
	"gorm.io/gorm"
)

type ProductTable struct {
	gorm.Model
	Id uint64
	Name string
	Des string
	Price float64
}

func (t ProductTable) TableName() string  {
	return "tb_product"
}


func (t ProductTable) QueryOneById(id int64) (*ProductTable, error)  {
	result := services.GetSqlConn().First(&t, id)
	return &t, result.Error
}

func (t ProductTable) InsertOne() (uint64, error)  {
	result := services.GetSqlConn().Create(&t)
	return t.Id, result.Error
}

