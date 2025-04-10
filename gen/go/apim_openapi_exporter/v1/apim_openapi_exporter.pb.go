// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: apim_openapi_exporter/v1/apim_openapi_exporter.proto

package apim_openapi_exporter

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/voikin/apim-proto/gen/go/shared/v1"
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

type BuildOpenAPISpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiGraph *v1.APIGraph `protobuf:"bytes,1,opt,name=api_graph,json=apiGraph,proto3" json:"api_graph,omitempty"`
}

func (x *BuildOpenAPISpecRequest) Reset() {
	*x = BuildOpenAPISpecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildOpenAPISpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildOpenAPISpecRequest) ProtoMessage() {}

func (x *BuildOpenAPISpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildOpenAPISpecRequest.ProtoReflect.Descriptor instead.
func (*BuildOpenAPISpecRequest) Descriptor() ([]byte, []int) {
	return file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescGZIP(), []int{0}
}

func (x *BuildOpenAPISpecRequest) GetApiGraph() *v1.APIGraph {
	if x != nil {
		return x.ApiGraph
	}
	return nil
}

type BuildOpenAPISpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecJson string `protobuf:"bytes,1,opt,name=spec_json,json=specJson,proto3" json:"spec_json,omitempty"`
}

func (x *BuildOpenAPISpecResponse) Reset() {
	*x = BuildOpenAPISpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildOpenAPISpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildOpenAPISpecResponse) ProtoMessage() {}

func (x *BuildOpenAPISpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildOpenAPISpecResponse.ProtoReflect.Descriptor instead.
func (*BuildOpenAPISpecResponse) Descriptor() ([]byte, []int) {
	return file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescGZIP(), []int{1}
}

func (x *BuildOpenAPISpecResponse) GetSpecJson() string {
	if x != nil {
		return x.SpecJson
	}
	return ""
}

var File_apim_openapi_exporter_v1_apim_openapi_exporter_proto protoreflect.FileDescriptor

var file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x5f,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x17, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x50, 0x49, 0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22, 0x37, 0x0a, 0x18, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63,
	0x4a, 0x73, 0x6f, 0x6e, 0x32, 0xcc, 0x01, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xb1, 0x01, 0x0a, 0x10, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x92, 0x41, 0x0e,
	0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x20, 0x53, 0x70, 0x65, 0x63, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x42, 0x76, 0x92, 0x41, 0x1f, 0x12, 0x1d, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e,
	0x41, 0x50, 0x49, 0x20, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49,
	0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x69, 0x6b, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescOnce sync.Once
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescData = file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDesc
)

func file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescGZIP() []byte {
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescOnce.Do(func() {
		file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescData)
	})
	return file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDescData
}

var file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_goTypes = []any{
	(*BuildOpenAPISpecRequest)(nil),  // 0: apim_openapi_exporter.v1.BuildOpenAPISpecRequest
	(*BuildOpenAPISpecResponse)(nil), // 1: apim_openapi_exporter.v1.BuildOpenAPISpecResponse
	(*v1.APIGraph)(nil),              // 2: shared.v1.APIGraph
}
var file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_depIdxs = []int32{
	2, // 0: apim_openapi_exporter.v1.BuildOpenAPISpecRequest.api_graph:type_name -> shared.v1.APIGraph
	0, // 1: apim_openapi_exporter.v1.OpenAPIExporterService.BuildOpenAPISpec:input_type -> apim_openapi_exporter.v1.BuildOpenAPISpecRequest
	1, // 2: apim_openapi_exporter.v1.OpenAPIExporterService.BuildOpenAPISpec:output_type -> apim_openapi_exporter.v1.BuildOpenAPISpecResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_init() }
func file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_init() {
	if File_apim_openapi_exporter_v1_apim_openapi_exporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BuildOpenAPISpecRequest); i {
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
		file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BuildOpenAPISpecResponse); i {
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
			RawDescriptor: file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_goTypes,
		DependencyIndexes: file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_depIdxs,
		MessageInfos:      file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_msgTypes,
	}.Build()
	File_apim_openapi_exporter_v1_apim_openapi_exporter_proto = out.File
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_rawDesc = nil
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_goTypes = nil
	file_apim_openapi_exporter_v1_apim_openapi_exporter_proto_depIdxs = nil
}
