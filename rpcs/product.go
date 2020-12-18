package rpcs

import (
	"context"
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/model"
)

var _ business.ProductServiceServer = (*ProductServiceServerImpl)(nil)

type ProductServiceServerImpl struct {
}

func (p ProductServiceServerImpl) ProductCreate(ctx context.Context, req *business.Product_ProductCreateReq) (*business.CommonRsp, error) {
	_, err := model.ProductModel{}.CreateOne(req.Name, req.Des, float64(req.Price))
	if err != nil{
		r := business.CommonRsp{
			Ret: business.ErrorCode_DatabaseError,
			Msg: err.Error(),
		}
		return &r, nil
	}
	r := business.CommonRsp{
		Ret: business.ErrorCode_Success,
		Msg: "success",
	}
	return &r, nil
}

func (p ProductServiceServerImpl) ProductList(ctx context.Context, req *business.Product_ProductListReq) (*business.Product_ProductListRsp, error) {
	data, err := model.ProductModel{}.QueryByNameOrDesOrPrice(req.Query, req.PriceLow, req.PriceHigh)
	if err != nil{
		r := business.Product_ProductListRsp{
			Ret:  business.ErrorCode_DatabaseError,
			Msg:  err.Error(),
			Data: nil,
		}
		return &r, nil
	}
	r := business.Product_ProductListRsp{
		Ret:  business.ErrorCode_Success,
		Msg:  "success",
		Data: data,
	}
	return &r, nil
}

func (p ProductServiceServerImpl) ProductDelete(ctx context.Context, req *business.Product_ProductDeleteReq) (*business.CommonRsp, error) {
	err := model.ProductModel{}.DeleteOne(uint64(req.ProductId))
	if err != nil{
		r := business.CommonRsp{
			Ret: business.ErrorCode_DatabaseError,
			Msg: err.Error(),
		}
		return &r, nil
	}
	r := business.CommonRsp{
		Ret: business.ErrorCode_Success,
		Msg: "success",
	}
	return &r, nil
}

func (p ProductServiceServerImpl) ProductUpdate(ctx context.Context, req *business.Product_ProductUpdateReq) (*business.CommonRsp, error) {
	err := model.ProductModel{}.UpdateOne(uint64(req.ProductId), req.Name, req.Des, float64(req.Price))
	if err != nil{
		r := business.CommonRsp{
			Ret: business.ErrorCode_DatabaseError,
			Msg: err.Error(),
		}
		return &r, nil
	}
	r := business.CommonRsp{
		Ret: business.ErrorCode_Success,
		Msg: "success",
	}
	return &r, nil
}
