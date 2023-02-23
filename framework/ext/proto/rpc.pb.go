// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc.proto

package proto

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

type ReqHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`          // 协议版本
	ReqId   uint32 `protobuf:"varint,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"` // 请求id
	Api     string `protobuf:"bytes,3,opt,name=api,proto3" json:"api,omitempty"`                   // 请求接口名
	Timeout uint32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`          // 超时时间
}

func (x *ReqHead) Reset() {
	*x = ReqHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqHead) ProtoMessage() {}

func (x *ReqHead) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqHead.ProtoReflect.Descriptor instead.
func (*ReqHead) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ReqHead) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ReqHead) GetReqId() uint32 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *ReqHead) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ReqHead) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type RspHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`          // 协议版本
	ReqId   uint32 `protobuf:"varint,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"` // 请求id
	Ret     int32  `protobuf:"varint,3,opt,name=ret,proto3" json:"ret,omitempty"`                  // 接口返回值，0:成功；1～1000框架错误码；其它为业务错误码
	Msg     []byte `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`                   // 接口返回值描述，可以是错误描述；也可以是其它
}

func (x *RspHead) Reset() {
	*x = RspHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspHead) ProtoMessage() {}

func (x *RspHead) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspHead.ProtoReflect.Descriptor instead.
func (*RspHead) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RspHead) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RspHead) GetReqId() uint32 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *RspHead) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *RspHead) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type RpcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"` // 协议头
	Body []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"` // 协议体
}

func (x *RpcReq) Reset() {
	*x = RpcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcReq) ProtoMessage() {}

func (x *RpcReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcReq.ProtoReflect.Descriptor instead.
func (*RpcReq) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RpcReq) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *RpcReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type RpcRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *RspHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"` // 协议头
	Body []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"` // 协议体
}

func (x *RpcRsp) Reset() {
	*x = RpcRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRsp) ProtoMessage() {}

func (x *RpcRsp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRsp.ProtoReflect.Descriptor instead.
func (*RpcRsp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *RpcRsp) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *RpcRsp) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x66, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x5e, 0x0a, 0x07, 0x52, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x40, 0x0a, 0x06, 0x52, 0x70,
	0x63, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65,
	0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x40, 0x0a, 0x06,
	0x52, 0x70, 0x63, 0x52, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x73, 0x70,
	0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_proto_goTypes = []interface{}{
	(*ReqHead)(nil), // 0: proto.ReqHead
	(*RspHead)(nil), // 1: proto.RspHead
	(*RpcReq)(nil),  // 2: proto.RpcReq
	(*RpcRsp)(nil),  // 3: proto.RpcRsp
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: proto.RpcReq.head:type_name -> proto.ReqHead
	1, // 1: proto.RpcRsp.head:type_name -> proto.RspHead
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqHead); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspHead); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcReq); i {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcRsp); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
