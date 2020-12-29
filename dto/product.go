package dto

import (
	"errors"
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/dao/table"
	"reflect"
)

type ProductDto struct {
}

func (dto *ProductDto) TableToInfo(t interface{}) (interface{}, error) {
	if reflect.TypeOf(t) != reflect.TypeOf(table.ProductTable{}) {
		return nil, errors.New("wrong table type")
	}

	return &business.Product_ProductListRsp_ProductInfo{
		ProductId: int64(t.(table.ProductTable).Id),
		Name:      t.(table.ProductTable).Name,
		Des:       t.(table.ProductTable).Des,
		Price:     float32(t.(table.ProductTable).Price),
	}, nil
}
