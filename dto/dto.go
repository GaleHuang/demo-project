package dto

import (
	"errors"
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/dao/table"
)

type CommonDtoFactory struct {
}

func NewDtoFactory() *CommonDtoFactory {
	return &CommonDtoFactory{}
}

func (f *CommonDtoFactory) CreateDto(name string) (interface{}, error) {
	if name == "product"{
		return &productDto{}, nil
	}
	return nil, errors.New("dto not found")
}


type ProductDto struct {
}

func (dto ProductDto) TableToInfo(table *table.ProductTable) *business.Product_ProductListRsp_ProductInfo  {
	return &business.Product_ProductListRsp_ProductInfo{
		ProductId: int64(table.Id),
		Name:      table.Name,
		Des:       table.Des,
		Price:     float32(table.Price),
	}
}










