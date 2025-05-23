// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: shared/v1/application.proto

package shared

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_shared_v1_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ApplicationProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ApplicationId string                 `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Version       uint32                 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ApiGraph      *APIGraph              `protobuf:"bytes,5,opt,name=api_graph,json=apiGraph,proto3,oneof" json:"api_graph,omitempty"`
}

func (x *ApplicationProfile) Reset() {
	*x = ApplicationProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationProfile) ProtoMessage() {}

func (x *ApplicationProfile) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationProfile.ProtoReflect.Descriptor instead.
func (*ApplicationProfile) Descriptor() ([]byte, []int) {
	return file_shared_v1_application_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApplicationProfile) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *ApplicationProfile) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ApplicationProfile) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApplicationProfile) GetApiGraph() *APIGraph {
	if x != nil {
		return x.ApiGraph
	}
	return nil
}

var File_shared_v1_application_proto protoreflect.FileDescriptor

var file_shared_v1_application_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xe5, 0x01, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x47, 0x72, 0x61, 0x70, 0x68, 0x48, 0x00, 0x52,
	0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x61, 0x70, 0x68, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x69, 0x6b, 0x69, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_v1_application_proto_rawDescOnce sync.Once
	file_shared_v1_application_proto_rawDescData = file_shared_v1_application_proto_rawDesc
)

func file_shared_v1_application_proto_rawDescGZIP() []byte {
	file_shared_v1_application_proto_rawDescOnce.Do(func() {
		file_shared_v1_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_v1_application_proto_rawDescData)
	})
	return file_shared_v1_application_proto_rawDescData
}

var file_shared_v1_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_v1_application_proto_goTypes = []any{
	(*Application)(nil),           // 0: shared.v1.Application
	(*ApplicationProfile)(nil),    // 1: shared.v1.ApplicationProfile
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*APIGraph)(nil),              // 3: shared.v1.APIGraph
}
var file_shared_v1_application_proto_depIdxs = []int32{
	2, // 0: shared.v1.Application.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: shared.v1.ApplicationProfile.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: shared.v1.ApplicationProfile.api_graph:type_name -> shared.v1.APIGraph
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shared_v1_application_proto_init() }
func file_shared_v1_application_proto_init() {
	if File_shared_v1_application_proto != nil {
		return
	}
	file_shared_v1_api_graph_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_v1_application_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Application); i {
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
		file_shared_v1_application_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ApplicationProfile); i {
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
	file_shared_v1_application_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_v1_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_v1_application_proto_goTypes,
		DependencyIndexes: file_shared_v1_application_proto_depIdxs,
		MessageInfos:      file_shared_v1_application_proto_msgTypes,
	}.Build()
	File_shared_v1_application_proto = out.File
	file_shared_v1_application_proto_rawDesc = nil
	file_shared_v1_application_proto_goTypes = nil
	file_shared_v1_application_proto_depIdxs = nil
}
