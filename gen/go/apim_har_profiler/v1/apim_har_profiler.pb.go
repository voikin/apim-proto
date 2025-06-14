// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: apim_har_profiler/v1/apim_har_profiler.proto

package apim_har_profiler

import (
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

type BuildAPIGraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HarFiles []*v1.HARFileWithFlags `protobuf:"bytes,1,rep,name=har_files,json=harFiles,proto3" json:"har_files,omitempty"`
}

func (x *BuildAPIGraphRequest) Reset() {
	*x = BuildAPIGraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildAPIGraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildAPIGraphRequest) ProtoMessage() {}

func (x *BuildAPIGraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildAPIGraphRequest.ProtoReflect.Descriptor instead.
func (*BuildAPIGraphRequest) Descriptor() ([]byte, []int) {
	return file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescGZIP(), []int{0}
}

func (x *BuildAPIGraphRequest) GetHarFiles() []*v1.HARFileWithFlags {
	if x != nil {
		return x.HarFiles
	}
	return nil
}

type BuildAPIGraphResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Graph *v1.APIGraph `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (x *BuildAPIGraphResponse) Reset() {
	*x = BuildAPIGraphResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildAPIGraphResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildAPIGraphResponse) ProtoMessage() {}

func (x *BuildAPIGraphResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildAPIGraphResponse.ProtoReflect.Descriptor instead.
func (*BuildAPIGraphResponse) Descriptor() ([]byte, []int) {
	return file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescGZIP(), []int{1}
}

func (x *BuildAPIGraphResponse) GetGraph() *v1.APIGraph {
	if x != nil {
		return x.Graph
	}
	return nil
}

var File_apim_har_profiler_v1_apim_har_profiler_proto protoreflect.FileDescriptor

var file_apim_har_profiler_v1_apim_har_profiler_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x50, 0x0a, 0x14, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x50, 0x49, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x68, 0x61,
	0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x41, 0x52, 0x46, 0x69, 0x6c,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x08, 0x68, 0x61, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x15, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x50, 0x49,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x32, 0xb1, 0x01, 0x0a, 0x12, 0x48, 0x41, 0x52,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x50, 0x49, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x50,
	0x49, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x50, 0x49, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x92, 0x41, 0x0b, 0x0a,
	0x09, 0x41, 0x50, 0x49, 0x20, 0x67, 0x72, 0x61, 0x70, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x42, 0x6a, 0x92, 0x41,
	0x1b, 0x12, 0x19, 0x0a, 0x10, 0x48, 0x41, 0x52, 0x20, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x20, 0x41, 0x50, 0x49, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x69, 0x6b, 0x69, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x6d, 0x5f, 0x68, 0x61, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescOnce sync.Once
	file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescData = file_apim_har_profiler_v1_apim_har_profiler_proto_rawDesc
)

func file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescGZIP() []byte {
	file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescOnce.Do(func() {
		file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescData)
	})
	return file_apim_har_profiler_v1_apim_har_profiler_proto_rawDescData
}

var file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apim_har_profiler_v1_apim_har_profiler_proto_goTypes = []any{
	(*BuildAPIGraphRequest)(nil),  // 0: apim_har_profiler.v1.BuildAPIGraphRequest
	(*BuildAPIGraphResponse)(nil), // 1: apim_har_profiler.v1.BuildAPIGraphResponse
	(*v1.HARFileWithFlags)(nil),   // 2: shared.v1.HARFileWithFlags
	(*v1.APIGraph)(nil),           // 3: shared.v1.APIGraph
}
var file_apim_har_profiler_v1_apim_har_profiler_proto_depIdxs = []int32{
	2, // 0: apim_har_profiler.v1.BuildAPIGraphRequest.har_files:type_name -> shared.v1.HARFileWithFlags
	3, // 1: apim_har_profiler.v1.BuildAPIGraphResponse.graph:type_name -> shared.v1.APIGraph
	0, // 2: apim_har_profiler.v1.HARProfilerService.BuildAPIGraph:input_type -> apim_har_profiler.v1.BuildAPIGraphRequest
	1, // 3: apim_har_profiler.v1.HARProfilerService.BuildAPIGraph:output_type -> apim_har_profiler.v1.BuildAPIGraphResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apim_har_profiler_v1_apim_har_profiler_proto_init() }
func file_apim_har_profiler_v1_apim_har_profiler_proto_init() {
	if File_apim_har_profiler_v1_apim_har_profiler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BuildAPIGraphRequest); i {
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
		file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BuildAPIGraphResponse); i {
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
			RawDescriptor: file_apim_har_profiler_v1_apim_har_profiler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apim_har_profiler_v1_apim_har_profiler_proto_goTypes,
		DependencyIndexes: file_apim_har_profiler_v1_apim_har_profiler_proto_depIdxs,
		MessageInfos:      file_apim_har_profiler_v1_apim_har_profiler_proto_msgTypes,
	}.Build()
	File_apim_har_profiler_v1_apim_har_profiler_proto = out.File
	file_apim_har_profiler_v1_apim_har_profiler_proto_rawDesc = nil
	file_apim_har_profiler_v1_apim_har_profiler_proto_goTypes = nil
	file_apim_har_profiler_v1_apim_har_profiler_proto_depIdxs = nil
}
