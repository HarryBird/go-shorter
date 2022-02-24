// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: v1/gateway.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenURLReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenURLReply) Reset() {
	*x = ShortenURLReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLReply) ProtoMessage() {}

func (x *ShortenURLReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLReply.ProtoReflect.Descriptor instead.
func (*ShortenURLReply) Descriptor() ([]byte, []int) {
	return file_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenURLReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DecodeURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecodeURLRequest) Reset() {
	*x = DecodeURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeURLRequest) ProtoMessage() {}

func (x *DecodeURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeURLRequest.ProtoReflect.Descriptor instead.
func (*DecodeURLRequest) Descriptor() ([]byte, []int) {
	return file_v1_gateway_proto_rawDescGZIP(), []int{2}
}

type DecodeURLReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecodeURLReply) Reset() {
	*x = DecodeURLReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeURLReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeURLReply) ProtoMessage() {}

func (x *DecodeURLReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeURLReply.ProtoReflect.Descriptor instead.
func (*DecodeURLReply) Descriptor() ([]byte, []int) {
	return file_v1_gateway_proto_rawDescGZIP(), []int{3}
}

var File_v1_gateway_proto protoreflect.FileDescriptor

var file_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x6d, 0x6f, 0x77, 0x65, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x11, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xfa,
	0x42, 0x14, 0x72, 0x12, 0x32, 0x10, 0x5e, 0x28, 0x68, 0x74, 0x74, 0x70, 0x7c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x29, 0x3a, 0x2f, 0x2f, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0f, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa8, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x90, 0x01, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52,
	0x4c, 0x12, 0x33, 0x2e, 0x6d, 0x6f, 0x77, 0x65, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x6f, 0x77, 0x65, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x32, 0x2e, 0x6d, 0x6f, 0x77, 0x65, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x6f, 0x77, 0x65, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_gateway_proto_rawDescOnce sync.Once
	file_v1_gateway_proto_rawDescData = file_v1_gateway_proto_rawDesc
)

func file_v1_gateway_proto_rawDescGZIP() []byte {
	file_v1_gateway_proto_rawDescOnce.Do(func() {
		file_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_gateway_proto_rawDescData)
	})
	return file_v1_gateway_proto_rawDescData
}

var file_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_gateway_proto_goTypes = []interface{}{
	(*ShortenURLRequest)(nil), // 0: mowen.api.url_shorten.gateway.v1.ShortenURLRequest
	(*ShortenURLReply)(nil),   // 1: mowen.api.url_shorten.gateway.v1.ShortenURLReply
	(*DecodeURLRequest)(nil),  // 2: mowen.api.url_shorten.gateway.v1.DecodeURLRequest
	(*DecodeURLReply)(nil),    // 3: mowen.api.url_shorten.gateway.v1.DecodeURLReply
}
var file_v1_gateway_proto_depIdxs = []int32{
	0, // 0: mowen.api.url_shorten.gateway.v1.Gateway.ShortenURL:input_type -> mowen.api.url_shorten.gateway.v1.ShortenURLRequest
	2, // 1: mowen.api.url_shorten.gateway.v1.Gateway.DecodeURL:input_type -> mowen.api.url_shorten.gateway.v1.DecodeURLRequest
	1, // 2: mowen.api.url_shorten.gateway.v1.Gateway.ShortenURL:output_type -> mowen.api.url_shorten.gateway.v1.ShortenURLReply
	3, // 3: mowen.api.url_shorten.gateway.v1.Gateway.DecodeURL:output_type -> mowen.api.url_shorten.gateway.v1.DecodeURLReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_gateway_proto_init() }
func file_v1_gateway_proto_init() {
	if File_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLRequest); i {
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
		file_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLReply); i {
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
		file_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeURLRequest); i {
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
		file_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeURLReply); i {
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
			RawDescriptor: file_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_gateway_proto_goTypes,
		DependencyIndexes: file_v1_gateway_proto_depIdxs,
		MessageInfos:      file_v1_gateway_proto_msgTypes,
	}.Build()
	File_v1_gateway_proto = out.File
	file_v1_gateway_proto_rawDesc = nil
	file_v1_gateway_proto_goTypes = nil
	file_v1_gateway_proto_depIdxs = nil
}
