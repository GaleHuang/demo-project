// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: common.proto

package business

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type EnumBusinessServiceName int32

const (
	EnumBusinessServiceName_BusinessProductService EnumBusinessServiceName = 0
)

// Enum value maps for EnumBusinessServiceName.
var (
	EnumBusinessServiceName_name = map[int32]string{
		0: "BusinessProductService",
	}
	EnumBusinessServiceName_value = map[string]int32{
		"BusinessProductService": 0,
	}
)

func (x EnumBusinessServiceName) Enum() *EnumBusinessServiceName {
	p := new(EnumBusinessServiceName)
	*p = x
	return p
}

func (x EnumBusinessServiceName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumBusinessServiceName) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (EnumBusinessServiceName) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x EnumBusinessServiceName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumBusinessServiceName.Descriptor instead.
func (EnumBusinessServiceName) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type ErrorCode int32

const (
	// 操作成功
	ErrorCode_Success ErrorCode = 0
	// 通用错误码(1~10000)
	ErrorCode_UnknownError        ErrorCode = 1
	ErrorCode_ParamError          ErrorCode = 2
	ErrorCode_NoAuthError         ErrorCode = 3
	ErrorCode_LogicError          ErrorCode = 4
	ErrorCode_DatabaseError       ErrorCode = 5
	ErrorCode_CacheError          ErrorCode = 6
	ErrorCode_ResultNotFoundError ErrorCode = 7
	// 模块自定义错误码(对应EnumWeTestServiceName的编号+1, 作为自定义错误码的万位)
	// product
	ErrorCode_ProductNotExistsError ErrorCode = 10001
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:     "Success",
		1:     "UnknownError",
		2:     "ParamError",
		3:     "NoAuthError",
		4:     "LogicError",
		5:     "DatabaseError",
		6:     "CacheError",
		7:     "ResultNotFoundError",
		10001: "ProductNotExistsError",
	}
	ErrorCode_value = map[string]int32{
		"Success":               0,
		"UnknownError":          1,
		"ParamError":            2,
		"NoAuthError":           3,
		"LogicError":            4,
		"DatabaseError":         5,
		"CacheError":            6,
		"ResultNotFoundError":   7,
		"ProductNotExistsError": 10001,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type CommonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret ErrorCode `protobuf:"varint,1,opt,name=ret,proto3,enum=business.ErrorCode" json:"ret,omitempty"`
	Msg string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonRsp) Reset() {
	*x = CommonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRsp) ProtoMessage() {}

func (x *CommonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRsp.ProtoReflect.Descriptor instead.
func (*CommonRsp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonRsp) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_Success
}

func (x *CommonRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

var file_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*EnumBusinessServiceName)(nil),
		Field:         50006,
		Name:          "business.business_service_name",
		Tag:           "varint,50006,opt,name=business_service_name,enum=business.EnumBusinessServiceName",
		Filename:      "common.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional business.EnumBusinessServiceName business_service_name = 50006;
	E_BusinessServiceName = &file_common_proto_extTypes[0]
)

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x2a, 0x35, 0x0a, 0x17,
	0x45, 0x6e, 0x75, 0x6d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x10, 0x00, 0x2a, 0xb3, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x04, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f,
	0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x12, 0x1a, 0x0a,
	0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x91, 0x4e, 0x3a, 0x7b, 0x0a, 0x15, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd6, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x13,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(EnumBusinessServiceName)(0),        // 0: business.EnumBusinessServiceName
	(ErrorCode)(0),                      // 1: business.ErrorCode
	(*CommonRsp)(nil),                   // 2: business.CommonRsp
	(*EmptyReq)(nil),                    // 3: business.EmptyReq
	(*descriptorpb.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
}
var file_common_proto_depIdxs = []int32{
	1, // 0: business.CommonRsp.ret:type_name -> business.ErrorCode
	4, // 1: business.business_service_name:extendee -> google.protobuf.ServiceOptions
	0, // 2: business.business_service_name:type_name -> business.EnumBusinessServiceName
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRsp); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
		ExtensionInfos:    file_common_proto_extTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}