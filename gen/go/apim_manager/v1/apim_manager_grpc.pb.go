// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: apim_manager/v1/apim_manager.proto

package apim_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ManagerService_GenerateOpenAPISpecFromHAR_FullMethodName = "/apim_manager.v1.ManagerService/GenerateOpenAPISpecFromHAR"
)

// ManagerServiceClient is the client API for ManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerServiceClient interface {
	GenerateOpenAPISpecFromHAR(ctx context.Context, in *GenerateOpenAPISpecFromHARRequest, opts ...grpc.CallOption) (*GenerateOpenAPISpecFromHARResponse, error)
}

type managerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerServiceClient(cc grpc.ClientConnInterface) ManagerServiceClient {
	return &managerServiceClient{cc}
}

func (c *managerServiceClient) GenerateOpenAPISpecFromHAR(ctx context.Context, in *GenerateOpenAPISpecFromHARRequest, opts ...grpc.CallOption) (*GenerateOpenAPISpecFromHARResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateOpenAPISpecFromHARResponse)
	err := c.cc.Invoke(ctx, ManagerService_GenerateOpenAPISpecFromHAR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServiceServer is the server API for ManagerService service.
// All implementations must embed UnimplementedManagerServiceServer
// for forward compatibility.
type ManagerServiceServer interface {
	GenerateOpenAPISpecFromHAR(context.Context, *GenerateOpenAPISpecFromHARRequest) (*GenerateOpenAPISpecFromHARResponse, error)
	mustEmbedUnimplementedManagerServiceServer()
}

// UnimplementedManagerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManagerServiceServer struct{}

func (UnimplementedManagerServiceServer) GenerateOpenAPISpecFromHAR(context.Context, *GenerateOpenAPISpecFromHARRequest) (*GenerateOpenAPISpecFromHARResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOpenAPISpecFromHAR not implemented")
}
func (UnimplementedManagerServiceServer) mustEmbedUnimplementedManagerServiceServer() {}
func (UnimplementedManagerServiceServer) testEmbeddedByValue()                        {}

// UnsafeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServiceServer will
// result in compilation errors.
type UnsafeManagerServiceServer interface {
	mustEmbedUnimplementedManagerServiceServer()
}

func RegisterManagerServiceServer(s grpc.ServiceRegistrar, srv ManagerServiceServer) {
	// If the following call pancis, it indicates UnimplementedManagerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManagerService_ServiceDesc, srv)
}

func _ManagerService_GenerateOpenAPISpecFromHAR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOpenAPISpecFromHARRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServiceServer).GenerateOpenAPISpecFromHAR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManagerService_GenerateOpenAPISpecFromHAR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServiceServer).GenerateOpenAPISpecFromHAR(ctx, req.(*GenerateOpenAPISpecFromHARRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerService_ServiceDesc is the grpc.ServiceDesc for ManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apim_manager.v1.ManagerService",
	HandlerType: (*ManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateOpenAPISpecFromHAR",
			Handler:    _ManagerService_GenerateOpenAPISpecFromHAR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apim_manager/v1/apim_manager.proto",
}
