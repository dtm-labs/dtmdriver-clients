// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: trans.proto

package pb

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

type AdjustInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	UserID int64 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *AdjustInfo) Reset() {
	*x = AdjustInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustInfo) ProtoMessage() {}

func (x *AdjustInfo) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustInfo.ProtoReflect.Descriptor instead.
func (*AdjustInfo) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{0}
}

func (x *AdjustInfo) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AdjustInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_trans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_trans_proto_rawDescGZIP(), []int{1}
}

var File_trans_proto protoreflect.FileDescriptor

var file_trans_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x69,
	0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x76, 0x63, 0x12, 0x2d, 0x0a, 0x07, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x49, 0x6e, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x41, 0x64,
	0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trans_proto_rawDescOnce sync.Once
	file_trans_proto_rawDescData = file_trans_proto_rawDesc
)

func file_trans_proto_rawDescGZIP() []byte {
	file_trans_proto_rawDescOnce.Do(func() {
		file_trans_proto_rawDescData = protoimpl.X.CompressGZIP(file_trans_proto_rawDescData)
	})
	return file_trans_proto_rawDescData
}

var file_trans_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trans_proto_goTypes = []interface{}{
	(*AdjustInfo)(nil), // 0: trans.AdjustInfo
	(*Response)(nil),   // 1: trans.Response
}
var file_trans_proto_depIdxs = []int32{
	0, // 0: trans.TransSvc.TransIn:input_type -> trans.AdjustInfo
	0, // 1: trans.TransSvc.TransOut:input_type -> trans.AdjustInfo
	1, // 2: trans.TransSvc.TransIn:output_type -> trans.Response
	1, // 3: trans.TransSvc.TransOut:output_type -> trans.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trans_proto_init() }
func file_trans_proto_init() {
	if File_trans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustInfo); i {
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
		file_trans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_trans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trans_proto_goTypes,
		DependencyIndexes: file_trans_proto_depIdxs,
		MessageInfos:      file_trans_proto_msgTypes,
	}.Build()
	File_trans_proto = out.File
	file_trans_proto_rawDesc = nil
	file_trans_proto_goTypes = nil
	file_trans_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransSvcClient is the client API for TransSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransSvcClient interface {
	TransIn(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error)
	TransOut(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error)
}

type transSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewTransSvcClient(cc grpc.ClientConnInterface) TransSvcClient {
	return &transSvcClient{cc}
}

func (c *transSvcClient) TransIn(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trans.TransSvc/TransIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transSvcClient) TransOut(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trans.TransSvc/TransOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransSvcServer is the server API for TransSvc service.
type TransSvcServer interface {
	TransIn(context.Context, *AdjustInfo) (*Response, error)
	TransOut(context.Context, *AdjustInfo) (*Response, error)
}

// UnimplementedTransSvcServer can be embedded to have forward compatible implementations.
type UnimplementedTransSvcServer struct {
}

func (*UnimplementedTransSvcServer) TransIn(context.Context, *AdjustInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransIn not implemented")
}
func (*UnimplementedTransSvcServer) TransOut(context.Context, *AdjustInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOut not implemented")
}

func RegisterTransSvcServer(s *grpc.Server, srv TransSvcServer) {
	s.RegisterService(&_TransSvc_serviceDesc, srv)
}

func _TransSvc_TransIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransSvcServer).TransIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trans.TransSvc/TransIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransSvcServer).TransIn(ctx, req.(*AdjustInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransSvc_TransOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransSvcServer).TransOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trans.TransSvc/TransOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransSvcServer).TransOut(ctx, req.(*AdjustInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trans.TransSvc",
	HandlerType: (*TransSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransIn",
			Handler:    _TransSvc_TransIn_Handler,
		},
		{
			MethodName: "TransOut",
			Handler:    _TransSvc_TransOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trans.proto",
}
