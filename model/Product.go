package model

import "github.com/galehuang/demo-project/dao/table"

type ProductModel struct {
}

func (ProductModel) CreateOne(name string, des string, price float64) (uint64, error)  {
	productDao := table.ProductTable{
		Name:  name,
		Des:   des,
		Price: price,
	}
	return productDao.InsertOne()
}

func (ProductModel) QueryByNameOrDes(query string) (*[]table.ProductTable, error) {
	return table.ProductTable{}.SearchByNameOrDes(query)
}

func (ProductModel) UpdateOne(id uint64, name string, des string, price float64) error  {
	productDao := table.ProductTable{
		Id:    id,
		Name:  name,
		Des:   des,
		Price: price,
	}
	return productDao.UpdateOne()
}

func (ProductModel) DeleteOne(id uint64) error  {
	return table.ProductTable{}.DeleteOne(id)
}


