package table

import "gorm.io/gorm"

type ProductTable struct {
	gorm.Model
	Id uint64
	Name string
	Des string
	Price float64
}

func (t ProductTable) Database()  {

}

func (t ProductTable) Table()  {

}

func (t ProductTable) QueryOneById(id int64) (*ProductTable, error)  {
	return
}
