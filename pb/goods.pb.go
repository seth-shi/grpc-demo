//adf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pb/goods.proto

package pb

import (
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

type GoodsShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GoodsShowRequest) Reset() {
	*x = GoodsShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsShowRequest) ProtoMessage() {}

func (x *GoodsShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsShowRequest.ProtoReflect.Descriptor instead.
func (*GoodsShowRequest) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsShowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoodsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Amount float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GoodsData) Reset() {
	*x = GoodsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsData) ProtoMessage() {}

func (x *GoodsData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsData.ProtoReflect.Descriptor instead.
func (*GoodsData) Descriptor() ([]byte, []int) {
	return file_pb_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsData) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_pb_goods_proto protoreflect.FileDescriptor

var file_pb_goods_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x68, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x34, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x68,
	0x6f, 0x77, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x68, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x74, 0x68, 0x2d, 0x73, 0x68, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_goods_proto_rawDescOnce sync.Once
	file_pb_goods_proto_rawDescData = file_pb_goods_proto_rawDesc
)

func file_pb_goods_proto_rawDescGZIP() []byte {
	file_pb_goods_proto_rawDescOnce.Do(func() {
		file_pb_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_goods_proto_rawDescData)
	})
	return file_pb_goods_proto_rawDescData
}

var file_pb_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_goods_proto_goTypes = []interface{}{
	(*GoodsShowRequest)(nil), // 0: pb.GoodsShowRequest
	(*GoodsData)(nil),        // 1: pb.GoodsData
}
var file_pb_goods_proto_depIdxs = []int32{
	0, // 0: pb.Goods.Show:input_type -> pb.GoodsShowRequest
	1, // 1: pb.Goods.Show:output_type -> pb.GoodsData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_goods_proto_init() }
func file_pb_goods_proto_init() {
	if File_pb_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsShowRequest); i {
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
		file_pb_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsData); i {
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
			RawDescriptor: file_pb_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_goods_proto_goTypes,
		DependencyIndexes: file_pb_goods_proto_depIdxs,
		MessageInfos:      file_pb_goods_proto_msgTypes,
	}.Build()
	File_pb_goods_proto = out.File
	file_pb_goods_proto_rawDesc = nil
	file_pb_goods_proto_goTypes = nil
	file_pb_goods_proto_depIdxs = nil
}
