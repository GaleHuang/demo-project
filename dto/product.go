package dto

import (
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/dao/table"
)

type productDto struct {
}

func (d productDto) TableToInfo(table *table.ProductTable) *business.Product_ProductListRsp_ProductInfo  {
	return &business.Product_ProductListRsp_ProductInfo{
		ProductId: int64(table.Id),
		Name:      table.Name,
		Des:       table.Des,
		Price:     float32(table.Price),
	}
}
