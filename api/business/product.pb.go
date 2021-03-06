// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: product.proto

package business

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

type Product_ProductCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`     // 产品名称
	Des   string  `protobuf:"bytes,2,opt,name=des,proto3" json:"des,omitempty"`       // 产品描述
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"` // 产品价格
}

func (x *Product_ProductCreateReq) Reset() {
	*x = Product_ProductCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductCreateReq) ProtoMessage() {}

func (x *Product_ProductCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductCreateReq.ProtoReflect.Descriptor instead.
func (*Product_ProductCreateReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Product_ProductCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product_ProductCreateReq) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *Product_ProductCreateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Product_ProductListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     string  `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`                            // 查询（名称或描述）
	PriceLow  float32 `protobuf:"fixed32,2,opt,name=price_low,json=priceLow,proto3" json:"price_low,omitempty"`    // 价格下限
	PriceHigh float32 `protobuf:"fixed32,3,opt,name=price_high,json=priceHigh,proto3" json:"price_high,omitempty"` // 价格上限
}

func (x *Product_ProductListReq) Reset() {
	*x = Product_ProductListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductListReq) ProtoMessage() {}

func (x *Product_ProductListReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductListReq.ProtoReflect.Descriptor instead.
func (*Product_ProductListReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Product_ProductListReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Product_ProductListReq) GetPriceLow() float32 {
	if x != nil {
		return x.PriceLow
	}
	return 0
}

func (x *Product_ProductListReq) GetPriceHigh() float32 {
	if x != nil {
		return x.PriceHigh
	}
	return 0
}

type Product_ProductListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret  ErrorCode                               `protobuf:"varint,1,opt,name=ret,proto3,enum=business.ErrorCode" json:"ret,omitempty"`
	Msg  string                                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *Product_ProductListRsp_ProductListData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Product_ProductListRsp) Reset() {
	*x = Product_ProductListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductListRsp) ProtoMessage() {}

func (x *Product_ProductListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductListRsp.ProtoReflect.Descriptor instead.
func (*Product_ProductListRsp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Product_ProductListRsp) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_Success
}

func (x *Product_ProductListRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Product_ProductListRsp) GetData() *Product_ProductListRsp_ProductListData {
	if x != nil {
		return x.Data
	}
	return nil
}

type Product_ProductDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *Product_ProductDeleteReq) Reset() {
	*x = Product_ProductDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductDeleteReq) ProtoMessage() {}

func (x *Product_ProductDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductDeleteReq.ProtoReflect.Descriptor instead.
func (*Product_ProductDeleteReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Product_ProductDeleteReq) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type Product_ProductUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64   `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Des       string  `protobuf:"bytes,3,opt,name=des,proto3" json:"des,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product_ProductUpdateReq) Reset() {
	*x = Product_ProductUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductUpdateReq) ProtoMessage() {}

func (x *Product_ProductUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductUpdateReq.ProtoReflect.Descriptor instead.
func (*Product_ProductUpdateReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Product_ProductUpdateReq) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product_ProductUpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product_ProductUpdateReq) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *Product_ProductUpdateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Product_ProductListRsp_ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64   `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Des       string  `protobuf:"bytes,3,opt,name=des,proto3" json:"des,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product_ProductListRsp_ProductInfo) Reset() {
	*x = Product_ProductListRsp_ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductListRsp_ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductListRsp_ProductInfo) ProtoMessage() {}

func (x *Product_ProductListRsp_ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductListRsp_ProductInfo.ProtoReflect.Descriptor instead.
func (*Product_ProductListRsp_ProductInfo) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Product_ProductListRsp_ProductInfo) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product_ProductListRsp_ProductInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product_ProductListRsp_ProductInfo) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *Product_ProductListRsp_ProductInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Product_ProductListRsp_ProductListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductList []*Product_ProductListRsp_ProductInfo `protobuf:"bytes,1,rep,name=product_list,json=productList,proto3" json:"product_list,omitempty"`
}

func (x *Product_ProductListRsp_ProductListData) Reset() {
	*x = Product_ProductListRsp_ProductListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product_ProductListRsp_ProductListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product_ProductListRsp_ProductListData) ProtoMessage() {}

func (x *Product_ProductListRsp_ProductListData) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product_ProductListRsp_ProductListData.ProtoReflect.Descriptor instead.
func (*Product_ProductListRsp_ProductListData) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *Product_ProductListRsp_ProductListData) GetProductList() []*Product_ProductListRsp_ProductInfo {
	if x != nil {
		return x.ProductList
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x05, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x1a, 0x4e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x1a, 0x62, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x48, 0x69, 0x67, 0x68, 0x1a, 0xdd, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x72, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x73, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x68, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x1a, 0x62, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x31, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x6d, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0xc7, 0x02, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x1a, 0x04, 0xb0,
	0xb5, 0x18, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_product_proto_goTypes = []interface{}{
	(*Product)(nil),                                // 0: business.Product
	(*Product_ProductCreateReq)(nil),               // 1: business.Product.ProductCreateReq
	(*Product_ProductListReq)(nil),                 // 2: business.Product.ProductListReq
	(*Product_ProductListRsp)(nil),                 // 3: business.Product.ProductListRsp
	(*Product_ProductDeleteReq)(nil),               // 4: business.Product.ProductDeleteReq
	(*Product_ProductUpdateReq)(nil),               // 5: business.Product.ProductUpdateReq
	(*Product_ProductListRsp_ProductInfo)(nil),     // 6: business.Product.ProductListRsp.ProductInfo
	(*Product_ProductListRsp_ProductListData)(nil), // 7: business.Product.ProductListRsp.ProductListData
	(ErrorCode)(0),                                 // 8: business.ErrorCode
	(*CommonRsp)(nil),                              // 9: business.CommonRsp
}
var file_product_proto_depIdxs = []int32{
	8, // 0: business.Product.ProductListRsp.ret:type_name -> business.ErrorCode
	7, // 1: business.Product.ProductListRsp.data:type_name -> business.Product.ProductListRsp.ProductListData
	6, // 2: business.Product.ProductListRsp.ProductListData.product_list:type_name -> business.Product.ProductListRsp.ProductInfo
	1, // 3: business.ProductService.ProductCreate:input_type -> business.Product.ProductCreateReq
	2, // 4: business.ProductService.ProductList:input_type -> business.Product.ProductListReq
	4, // 5: business.ProductService.ProductDelete:input_type -> business.Product.ProductDeleteReq
	5, // 6: business.ProductService.ProductUpdate:input_type -> business.Product.ProductUpdateReq
	9, // 7: business.ProductService.ProductCreate:output_type -> business.CommonRsp
	3, // 8: business.ProductService.ProductList:output_type -> business.Product.ProductListRsp
	9, // 9: business.ProductService.ProductDelete:output_type -> business.CommonRsp
	9, // 10: business.ProductService.ProductUpdate:output_type -> business.CommonRsp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductListRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductDeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductListRsp_ProductInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product_ProductListRsp_ProductListData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 添加产品
	ProductCreate(ctx context.Context, in *Product_ProductCreateReq, opts ...grpc.CallOption) (*CommonRsp, error)
	// 查询产品
	ProductList(ctx context.Context, in *Product_ProductListReq, opts ...grpc.CallOption) (*Product_ProductListRsp, error)
	// 删除产品
	ProductDelete(ctx context.Context, in *Product_ProductDeleteReq, opts ...grpc.CallOption) (*CommonRsp, error)
	// 更新产品
	ProductUpdate(ctx context.Context, in *Product_ProductUpdateReq, opts ...grpc.CallOption) (*CommonRsp, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) ProductCreate(ctx context.Context, in *Product_ProductCreateReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, "/business.ProductService/ProductCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ProductList(ctx context.Context, in *Product_ProductListReq, opts ...grpc.CallOption) (*Product_ProductListRsp, error) {
	out := new(Product_ProductListRsp)
	err := c.cc.Invoke(ctx, "/business.ProductService/ProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ProductDelete(ctx context.Context, in *Product_ProductDeleteReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, "/business.ProductService/ProductDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ProductUpdate(ctx context.Context, in *Product_ProductUpdateReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, "/business.ProductService/ProductUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// 添加产品
	ProductCreate(context.Context, *Product_ProductCreateReq) (*CommonRsp, error)
	// 查询产品
	ProductList(context.Context, *Product_ProductListReq) (*Product_ProductListRsp, error)
	// 删除产品
	ProductDelete(context.Context, *Product_ProductDeleteReq) (*CommonRsp, error)
	// 更新产品
	ProductUpdate(context.Context, *Product_ProductUpdateReq) (*CommonRsp, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) ProductCreate(context.Context, *Product_ProductCreateReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductCreate not implemented")
}
func (*UnimplementedProductServiceServer) ProductList(context.Context, *Product_ProductListReq) (*Product_ProductListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (*UnimplementedProductServiceServer) ProductDelete(context.Context, *Product_ProductDeleteReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDelete not implemented")
}
func (*UnimplementedProductServiceServer) ProductUpdate(context.Context, *Product_ProductUpdateReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductUpdate not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_ProductCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product_ProductCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.ProductService/ProductCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductCreate(ctx, req.(*Product_ProductCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product_ProductListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.ProductService/ProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductList(ctx, req.(*Product_ProductListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product_ProductDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.ProductService/ProductDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductDelete(ctx, req.(*Product_ProductDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ProductUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product_ProductUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.ProductService/ProductUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductUpdate(ctx, req.(*Product_ProductUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "business.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductCreate",
			Handler:    _ProductService_ProductCreate_Handler,
		},
		{
			MethodName: "ProductList",
			Handler:    _ProductService_ProductList_Handler,
		},
		{
			MethodName: "ProductDelete",
			Handler:    _ProductService_ProductDelete_Handler,
		},
		{
			MethodName: "ProductUpdate",
			Handler:    _ProductService_ProductUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
