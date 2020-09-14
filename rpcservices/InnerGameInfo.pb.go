// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: InnerGameInfo.proto

package rpcservices

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

//接口请求入参
type InnerGameInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameCode string `protobuf:"bytes,1,opt,name=gameCode,proto3" json:"gameCode,omitempty"`
	Flag     string `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *InnerGameInfoRequest) Reset() {
	*x = InnerGameInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InnerGameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerGameInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerGameInfoRequest) ProtoMessage() {}

func (x *InnerGameInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_InnerGameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerGameInfoRequest.ProtoReflect.Descriptor instead.
func (*InnerGameInfoRequest) Descriptor() ([]byte, []int) {
	return file_InnerGameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InnerGameInfoRequest) GetGameCode() string {
	if x != nil {
		return x.GameCode
	}
	return ""
}

func (x *InnerGameInfoRequest) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

//接口返回出参
type InnerGameInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameInfoInnerVo []*GameInfoInnerVo `protobuf:"bytes,1,rep,name=gameInfoInnerVo,proto3" json:"gameInfoInnerVo,omitempty"`
}

func (x *InnerGameInfoResponse) Reset() {
	*x = InnerGameInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InnerGameInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerGameInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerGameInfoResponse) ProtoMessage() {}

func (x *InnerGameInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_InnerGameInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerGameInfoResponse.ProtoReflect.Descriptor instead.
func (*InnerGameInfoResponse) Descriptor() ([]byte, []int) {
	return file_InnerGameInfo_proto_rawDescGZIP(), []int{1}
}

func (x *InnerGameInfoResponse) GetGameInfoInnerVo() []*GameInfoInnerVo {
	if x != nil {
		return x.GameInfoInnerVo
	}
	return nil
}

var File_InnerGameInfo_proto protoreflect.FileDescriptor

var file_InnerGameInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x14, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x22, 0x59, 0x0a, 0x15, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x52, 0x0f, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x32, 0x64, 0x0a,
	0x14, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InnerGameInfo_proto_rawDescOnce sync.Once
	file_InnerGameInfo_proto_rawDescData = file_InnerGameInfo_proto_rawDesc
)

func file_InnerGameInfo_proto_rawDescGZIP() []byte {
	file_InnerGameInfo_proto_rawDescOnce.Do(func() {
		file_InnerGameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InnerGameInfo_proto_rawDescData)
	})
	return file_InnerGameInfo_proto_rawDescData
}

var file_InnerGameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_InnerGameInfo_proto_goTypes = []interface{}{
	(*InnerGameInfoRequest)(nil),  // 0: proto.InnerGameInfoRequest
	(*InnerGameInfoResponse)(nil), // 1: proto.InnerGameInfoResponse
	(*GameInfoInnerVo)(nil),       // 2: proto.GameInfoInnerVo
}
var file_InnerGameInfo_proto_depIdxs = []int32{
	2, // 0: proto.InnerGameInfoResponse.gameInfoInnerVo:type_name -> proto.GameInfoInnerVo
	0, // 1: proto.InnerGameInfoService.InnerGameInfo:input_type -> proto.InnerGameInfoRequest
	1, // 2: proto.InnerGameInfoService.InnerGameInfo:output_type -> proto.InnerGameInfoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InnerGameInfo_proto_init() }
func file_InnerGameInfo_proto_init() {
	if File_InnerGameInfo_proto != nil {
		return
	}
	file_Models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InnerGameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerGameInfoRequest); i {
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
		file_InnerGameInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerGameInfoResponse); i {
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
			RawDescriptor: file_InnerGameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_InnerGameInfo_proto_goTypes,
		DependencyIndexes: file_InnerGameInfo_proto_depIdxs,
		MessageInfos:      file_InnerGameInfo_proto_msgTypes,
	}.Build()
	File_InnerGameInfo_proto = out.File
	file_InnerGameInfo_proto_rawDesc = nil
	file_InnerGameInfo_proto_goTypes = nil
	file_InnerGameInfo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InnerGameInfoServiceClient is the client API for InnerGameInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InnerGameInfoServiceClient interface {
	InnerGameInfo(ctx context.Context, in *InnerGameInfoRequest, opts ...grpc.CallOption) (*InnerGameInfoResponse, error)
}

type innerGameInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInnerGameInfoServiceClient(cc grpc.ClientConnInterface) InnerGameInfoServiceClient {
	return &innerGameInfoServiceClient{cc}
}

func (c *innerGameInfoServiceClient) InnerGameInfo(ctx context.Context, in *InnerGameInfoRequest, opts ...grpc.CallOption) (*InnerGameInfoResponse, error) {
	out := new(InnerGameInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.InnerGameInfoService/InnerGameInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerGameInfoServiceServer is the server API for InnerGameInfoService service.
type InnerGameInfoServiceServer interface {
	InnerGameInfo(context.Context, *InnerGameInfoRequest) (*InnerGameInfoResponse, error)
}

// UnimplementedInnerGameInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInnerGameInfoServiceServer struct {
}

func (*UnimplementedInnerGameInfoServiceServer) InnerGameInfo(context.Context, *InnerGameInfoRequest) (*InnerGameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InnerGameInfo not implemented")
}

func RegisterInnerGameInfoServiceServer(s *grpc.Server, srv InnerGameInfoServiceServer) {
	s.RegisterService(&_InnerGameInfoService_serviceDesc, srv)
}

func _InnerGameInfoService_InnerGameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerGameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerGameInfoServiceServer).InnerGameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InnerGameInfoService/InnerGameInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerGameInfoServiceServer).InnerGameInfo(ctx, req.(*InnerGameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InnerGameInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InnerGameInfoService",
	HandlerType: (*InnerGameInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InnerGameInfo",
			Handler:    _InnerGameInfoService_InnerGameInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "InnerGameInfo.proto",
}
