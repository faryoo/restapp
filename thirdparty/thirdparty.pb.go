// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: thirdparty.proto

package thirdparty

import (
	context "context"
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

type TemplateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Touser      string               `protobuf:"bytes,1,opt,name=touser,proto3" json:"touser,omitempty"`
	TemplateId  string               `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Url         string               `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Color       string               `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Data        map[string]*MapVaule `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Miniprogram *MiniProgram         `protobuf:"bytes,6,opt,name=miniprogram,proto3" json:"miniprogram,omitempty"`
}

func (x *TemplateReq) Reset() {
	*x = TemplateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateReq) ProtoMessage() {}

func (x *TemplateReq) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateReq.ProtoReflect.Descriptor instead.
func (*TemplateReq) Descriptor() ([]byte, []int) {
	return file_thirdparty_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateReq) GetTouser() string {
	if x != nil {
		return x.Touser
	}
	return ""
}

func (x *TemplateReq) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *TemplateReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TemplateReq) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TemplateReq) GetData() map[string]*MapVaule {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TemplateReq) GetMiniprogram() *MiniProgram {
	if x != nil {
		return x.Miniprogram
	}
	return nil
}

type MapVaule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *MapVaule) Reset() {
	*x = MapVaule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapVaule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapVaule) ProtoMessage() {}

func (x *MapVaule) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapVaule.ProtoReflect.Descriptor instead.
func (*MapVaule) Descriptor() ([]byte, []int) {
	return file_thirdparty_proto_rawDescGZIP(), []int{1}
}

func (x *MapVaule) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MapVaule) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type MiniProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid    string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Pagepath string `protobuf:"bytes,2,opt,name=pagepath,proto3" json:"pagepath,omitempty"`
}

func (x *MiniProgram) Reset() {
	*x = MiniProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniProgram) ProtoMessage() {}

func (x *MiniProgram) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniProgram.ProtoReflect.Descriptor instead.
func (*MiniProgram) Descriptor() ([]byte, []int) {
	return file_thirdparty_proto_rawDescGZIP(), []int{2}
}

func (x *MiniProgram) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *MiniProgram) GetPagepath() string {
	if x != nil {
		return x.Pagepath
	}
	return ""
}

type TemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TemplateReply) Reset() {
	*x = TemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thirdparty_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateReply) ProtoMessage() {}

func (x *TemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_thirdparty_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateReply.ProtoReflect.Descriptor instead.
func (*TemplateReply) Descriptor() ([]byte, []int) {
	return file_thirdparty_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateReply) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

var File_thirdparty_proto protoreflect.FileDescriptor

var file_thirdparty_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0xaf,
	0x02, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x1a, 0x4d, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x4d, 0x61, 0x70,
	0x56, 0x61, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x69,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0x29, 0x0a, 0x0d, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0x99, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4d, 0x73, 0x67, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4d, 0x73, 0x67, 0x54, 0x6f, 0x43, 0x12, 0x17, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x19, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0e, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x42, 0x12, 0x17, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x61, 0x72, 0x79, 0x6f, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thirdparty_proto_rawDescOnce sync.Once
	file_thirdparty_proto_rawDescData = file_thirdparty_proto_rawDesc
)

func file_thirdparty_proto_rawDescGZIP() []byte {
	file_thirdparty_proto_rawDescOnce.Do(func() {
		file_thirdparty_proto_rawDescData = protoimpl.X.CompressGZIP(file_thirdparty_proto_rawDescData)
	})
	return file_thirdparty_proto_rawDescData
}

var file_thirdparty_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_thirdparty_proto_goTypes = []interface{}{
	(*TemplateReq)(nil),   // 0: thirdparty.TemplateReq
	(*MapVaule)(nil),      // 1: thirdparty.MapVaule
	(*MiniProgram)(nil),   // 2: thirdparty.MiniProgram
	(*TemplateReply)(nil), // 3: thirdparty.TemplateReply
	nil,                   // 4: thirdparty.TemplateReq.DataEntry
}
var file_thirdparty_proto_depIdxs = []int32{
	4, // 0: thirdparty.TemplateReq.data:type_name -> thirdparty.TemplateReq.DataEntry
	2, // 1: thirdparty.TemplateReq.miniprogram:type_name -> thirdparty.MiniProgram
	1, // 2: thirdparty.TemplateReq.DataEntry.value:type_name -> thirdparty.MapVaule
	0, // 3: thirdparty.TemplateMsg.templateMsgToC:input_type -> thirdparty.TemplateReq
	0, // 4: thirdparty.TemplateMsg.templateMsgToB:input_type -> thirdparty.TemplateReq
	3, // 5: thirdparty.TemplateMsg.templateMsgToC:output_type -> thirdparty.TemplateReply
	3, // 6: thirdparty.TemplateMsg.templateMsgToB:output_type -> thirdparty.TemplateReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_thirdparty_proto_init() }
func file_thirdparty_proto_init() {
	if File_thirdparty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thirdparty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateReq); i {
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
		file_thirdparty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapVaule); i {
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
		file_thirdparty_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiniProgram); i {
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
		file_thirdparty_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateReply); i {
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
			RawDescriptor: file_thirdparty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thirdparty_proto_goTypes,
		DependencyIndexes: file_thirdparty_proto_depIdxs,
		MessageInfos:      file_thirdparty_proto_msgTypes,
	}.Build()
	File_thirdparty_proto = out.File
	file_thirdparty_proto_rawDesc = nil
	file_thirdparty_proto_goTypes = nil
	file_thirdparty_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemplateMsgClient is the client API for TemplateMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateMsgClient interface {
	TemplateMsgToC(ctx context.Context, in *TemplateReq, opts ...grpc.CallOption) (*TemplateReply, error)
	TemplateMsgToB(ctx context.Context, in *TemplateReq, opts ...grpc.CallOption) (*TemplateReply, error)
}

type templateMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateMsgClient(cc grpc.ClientConnInterface) TemplateMsgClient {
	return &templateMsgClient{cc}
}

func (c *templateMsgClient) TemplateMsgToC(ctx context.Context, in *TemplateReq, opts ...grpc.CallOption) (*TemplateReply, error) {
	out := new(TemplateReply)
	err := c.cc.Invoke(ctx, "/thirdparty.TemplateMsg/templateMsgToC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateMsgClient) TemplateMsgToB(ctx context.Context, in *TemplateReq, opts ...grpc.CallOption) (*TemplateReply, error) {
	out := new(TemplateReply)
	err := c.cc.Invoke(ctx, "/thirdparty.TemplateMsg/templateMsgToB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateMsgServer is the server API for TemplateMsg service.
type TemplateMsgServer interface {
	TemplateMsgToC(context.Context, *TemplateReq) (*TemplateReply, error)
	TemplateMsgToB(context.Context, *TemplateReq) (*TemplateReply, error)
}

// UnimplementedTemplateMsgServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateMsgServer struct {
}

func (*UnimplementedTemplateMsgServer) TemplateMsgToC(context.Context, *TemplateReq) (*TemplateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateMsgToC not implemented")
}
func (*UnimplementedTemplateMsgServer) TemplateMsgToB(context.Context, *TemplateReq) (*TemplateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateMsgToB not implemented")
}

func RegisterTemplateMsgServer(s *grpc.Server, srv TemplateMsgServer) {
	s.RegisterService(&_TemplateMsg_serviceDesc, srv)
}

func _TemplateMsg_TemplateMsgToC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMsgServer).TemplateMsgToC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdparty.TemplateMsg/TemplateMsgToC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMsgServer).TemplateMsgToC(ctx, req.(*TemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateMsg_TemplateMsgToB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMsgServer).TemplateMsgToB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thirdparty.TemplateMsg/TemplateMsgToB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMsgServer).TemplateMsgToB(ctx, req.(*TemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemplateMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thirdparty.TemplateMsg",
	HandlerType: (*TemplateMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "templateMsgToC",
			Handler:    _TemplateMsg_TemplateMsgToC_Handler,
		},
		{
			MethodName: "templateMsgToB",
			Handler:    _TemplateMsg_TemplateMsgToB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thirdparty.proto",
}
