// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: apim_har_profiler/v1/apim_har_profiler.proto

/*
Package apim_har_profiler is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package apim_har_profiler

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_HARProfilerService_BuildAPIGraph_0(ctx context.Context, marshaler runtime.Marshaler, client HARProfilerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq BuildAPIGraphRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.BuildAPIGraph(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_HARProfilerService_BuildAPIGraph_0(ctx context.Context, marshaler runtime.Marshaler, server HARProfilerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq BuildAPIGraphRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.BuildAPIGraph(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterHARProfilerServiceHandlerServer registers the http handlers for service HARProfilerService to "mux".
// UnaryRPC     :call HARProfilerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterHARProfilerServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterHARProfilerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server HARProfilerServiceServer) error {
	mux.Handle(http.MethodPost, pattern_HARProfilerService_BuildAPIGraph_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/apim_har_profiler.v1.HARProfilerService/BuildAPIGraph", runtime.WithHTTPPathPattern("/api/v1/build-api-graph"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_HARProfilerService_BuildAPIGraph_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_HARProfilerService_BuildAPIGraph_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterHARProfilerServiceHandlerFromEndpoint is same as RegisterHARProfilerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterHARProfilerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterHARProfilerServiceHandler(ctx, mux, conn)
}

// RegisterHARProfilerServiceHandler registers the http handlers for service HARProfilerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterHARProfilerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterHARProfilerServiceHandlerClient(ctx, mux, NewHARProfilerServiceClient(conn))
}

// RegisterHARProfilerServiceHandlerClient registers the http handlers for service HARProfilerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "HARProfilerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "HARProfilerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "HARProfilerServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterHARProfilerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client HARProfilerServiceClient) error {
	mux.Handle(http.MethodPost, pattern_HARProfilerService_BuildAPIGraph_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/apim_har_profiler.v1.HARProfilerService/BuildAPIGraph", runtime.WithHTTPPathPattern("/api/v1/build-api-graph"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_HARProfilerService_BuildAPIGraph_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_HARProfilerService_BuildAPIGraph_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_HARProfilerService_BuildAPIGraph_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "build-api-graph"}, ""))
)

var (
	forward_HARProfilerService_BuildAPIGraph_0 = runtime.ForwardResponseMessage
)
