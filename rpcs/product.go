package rpcs

import (
	"context"
	"github.com/galehuang/demo-project/api/business"
)

var _ business.ProductServiceServer = (*ProductServiceServerImpl)(nil)

type ProductServiceServerImpl struct {
}

func (p ProductServiceServerImpl) ProductCreate(ctx context.Context, req *business.Product_ProductCreateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (p ProductServiceServerImpl) ProductList(ctx context.Context, req *business.Product_ProductListReq) (*business.Product_ProductListRsp, error) {
	panic("implement me")
}

func (p ProductServiceServerImpl) ProductDelete(ctx context.Context, req *business.Product_ProductDeleteReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (p ProductServiceServerImpl) ProductUpdate(ctx context.Context, req *business.Product_ProductUpdateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

