package model

import (
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/dao/table"
	"github.com/galehuang/demo-project/dto"
)

type ProductModel struct {
}

func (ProductModel) CreateOne(name string, des string, price float64) (uint64, error) {
	productDao := table.ProductTable{
		Name:  name,
		Des:   des,
		Price: price,
	}
	return productDao.InsertOne()
}

func (ProductModel) QueryByNameOrDesOrPrice(query string, priceLow float32, priceHigh float32) (*business.Product_ProductListRsp_ProductListData, error) {
	productList := make([]*business.Product_ProductListRsp_ProductInfo, 0)
	productDto, err := dto.NewDtoFactory().CreateDto("product")
	if err != nil {
		return nil, err
	}

	var result *[]table.ProductTable
	result, err = table.ProductTable{}.QueryByPriceOrNameOrDes(query, priceLow, priceHigh)
	if err != nil {
		return nil, err
	}

	for _, res := range *result {
		productInfo, err := productDto.TableToInfo(&res)
		if err != nil {
			return nil, err
		}
		productList = append(productList, productInfo.(*business.Product_ProductListRsp_ProductInfo))
	}
	data := &business.Product_ProductListRsp_ProductListData{ProductList: productList}
	return data, nil
}

func (ProductModel) UpdateOne(id uint64, name string, des string, price float64) error {
	productDao := table.ProductTable{
		Id:    id,
		Name:  name,
		Des:   des,
		Price: price,
	}
	return productDao.UpdateOne()
}

func (ProductModel) DeleteOne(id uint64) error {
	return table.ProductTable{}.DeleteOne(id)
}
