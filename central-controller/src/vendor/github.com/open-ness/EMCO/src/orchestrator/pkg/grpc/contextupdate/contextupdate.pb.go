//
// Copyright 2020 Intel Corporation, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: contextupdate.proto

package contextupdate

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

type ContextUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppContext string `protobuf:"bytes,1,opt,name=app_context,json=appContext,proto3" json:"app_context,omitempty"`
	IntentName string `protobuf:"bytes,2,opt,name=intent_name,json=intentName,proto3" json:"intent_name,omitempty"`
}

func (x *ContextUpdateRequest) Reset() {
	*x = ContextUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contextupdate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextUpdateRequest) ProtoMessage() {}

func (x *ContextUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contextupdate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextUpdateRequest.ProtoReflect.Descriptor instead.
func (*ContextUpdateRequest) Descriptor() ([]byte, []int) {
	return file_contextupdate_proto_rawDescGZIP(), []int{0}
}

func (x *ContextUpdateRequest) GetAppContext() string {
	if x != nil {
		return x.AppContext
	}
	return ""
}

func (x *ContextUpdateRequest) GetIntentName() string {
	if x != nil {
		return x.IntentName
	}
	return ""
}

type ContextUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppContextUpdated       bool   `protobuf:"varint,1,opt,name=app_context_updated,json=appContextUpdated,proto3" json:"app_context_updated,omitempty"`
	AppContextUpdateMessage string `protobuf:"bytes,2,opt,name=app_context_update_message,json=appContextUpdateMessage,proto3" json:"app_context_update_message,omitempty"`
}

func (x *ContextUpdateResponse) Reset() {
	*x = ContextUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contextupdate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextUpdateResponse) ProtoMessage() {}

func (x *ContextUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contextupdate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextUpdateResponse.ProtoReflect.Descriptor instead.
func (*ContextUpdateResponse) Descriptor() ([]byte, []int) {
	return file_contextupdate_proto_rawDescGZIP(), []int{1}
}

func (x *ContextUpdateResponse) GetAppContextUpdated() bool {
	if x != nil {
		return x.AppContextUpdated
	}
	return false
}

func (x *ContextUpdateResponse) GetAppContextUpdateMessage() string {
	if x != nil {
		return x.AppContextUpdateMessage
	}
	return ""
}

var File_contextupdate_proto protoreflect.FileDescriptor

var file_contextupdate_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x84, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x70, 0x70,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x61, 0x70, 0x70,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x54, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contextupdate_proto_rawDescOnce sync.Once
	file_contextupdate_proto_rawDescData = file_contextupdate_proto_rawDesc
)

func file_contextupdate_proto_rawDescGZIP() []byte {
	file_contextupdate_proto_rawDescOnce.Do(func() {
		file_contextupdate_proto_rawDescData = protoimpl.X.CompressGZIP(file_contextupdate_proto_rawDescData)
	})
	return file_contextupdate_proto_rawDescData
}

var file_contextupdate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contextupdate_proto_goTypes = []interface{}{
	(*ContextUpdateRequest)(nil),  // 0: ContextUpdateRequest
	(*ContextUpdateResponse)(nil), // 1: ContextUpdateResponse
}
var file_contextupdate_proto_depIdxs = []int32{
	0, // 0: contextupdate.UpdateAppContext:input_type -> ContextUpdateRequest
	1, // 1: contextupdate.UpdateAppContext:output_type -> ContextUpdateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contextupdate_proto_init() }
func file_contextupdate_proto_init() {
	if File_contextupdate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contextupdate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextUpdateRequest); i {
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
		file_contextupdate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextUpdateResponse); i {
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
			RawDescriptor: file_contextupdate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contextupdate_proto_goTypes,
		DependencyIndexes: file_contextupdate_proto_depIdxs,
		MessageInfos:      file_contextupdate_proto_msgTypes,
	}.Build()
	File_contextupdate_proto = out.File
	file_contextupdate_proto_rawDesc = nil
	file_contextupdate_proto_goTypes = nil
	file_contextupdate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContextupdateClient is the client API for Contextupdate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContextupdateClient interface {
	// Controllers
	UpdateAppContext(ctx context.Context, in *ContextUpdateRequest, opts ...grpc.CallOption) (*ContextUpdateResponse, error)
}

type contextupdateClient struct {
	cc grpc.ClientConnInterface
}

func NewContextupdateClient(cc grpc.ClientConnInterface) ContextupdateClient {
	return &contextupdateClient{cc}
}

func (c *contextupdateClient) UpdateAppContext(ctx context.Context, in *ContextUpdateRequest, opts ...grpc.CallOption) (*ContextUpdateResponse, error) {
	out := new(ContextUpdateResponse)
	err := c.cc.Invoke(ctx, "/contextupdate/UpdateAppContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextupdateServer is the server API for Contextupdate service.
type ContextupdateServer interface {
	// Controllers
	UpdateAppContext(context.Context, *ContextUpdateRequest) (*ContextUpdateResponse, error)
}

// UnimplementedContextupdateServer can be embedded to have forward compatible implementations.
type UnimplementedContextupdateServer struct {
}

func (*UnimplementedContextupdateServer) UpdateAppContext(context.Context, *ContextUpdateRequest) (*ContextUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppContext not implemented")
}

func RegisterContextupdateServer(s *grpc.Server, srv ContextupdateServer) {
	s.RegisterService(&_Contextupdate_serviceDesc, srv)
}

func _Contextupdate_UpdateAppContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextupdateServer).UpdateAppContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contextupdate/UpdateAppContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextupdateServer).UpdateAppContext(ctx, req.(*ContextUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contextupdate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contextupdate",
	HandlerType: (*ContextupdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAppContext",
			Handler:    _Contextupdate_UpdateAppContext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contextupdate.proto",
}
