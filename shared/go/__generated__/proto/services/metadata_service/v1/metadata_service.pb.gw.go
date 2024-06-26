// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: services/metadata_service/v1/metadata_service.proto

/*
Package metadata_servicev1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package metadata_servicev1

import (
	"context"
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
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_MetaDataService_UpdateIptcData_0(ctx context.Context, marshaler runtime.Marshaler, client MetaDataServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateIPTCDataRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateIptcData(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MetaDataService_UpdateIptcData_0(ctx context.Context, marshaler runtime.Marshaler, server MetaDataServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateIPTCDataRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateIptcData(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterMetaDataServiceHandlerServer registers the http handlers for service MetaDataService to "mux".
// UnaryRPC     :call MetaDataServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMetaDataServiceHandlerFromEndpoint instead.
func RegisterMetaDataServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MetaDataServiceServer) error {

	mux.Handle("POST", pattern_MetaDataService_UpdateIptcData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/metadata_service.v1.MetaDataService/UpdateIptcData", runtime.WithHTTPPathPattern("/v1/iptc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MetaDataService_UpdateIptcData_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MetaDataService_UpdateIptcData_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterMetaDataServiceHandlerFromEndpoint is same as RegisterMetaDataServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMetaDataServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterMetaDataServiceHandler(ctx, mux, conn)
}

// RegisterMetaDataServiceHandler registers the http handlers for service MetaDataService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMetaDataServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMetaDataServiceHandlerClient(ctx, mux, NewMetaDataServiceClient(conn))
}

// RegisterMetaDataServiceHandlerClient registers the http handlers for service MetaDataService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MetaDataServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MetaDataServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MetaDataServiceClient" to call the correct interceptors.
func RegisterMetaDataServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MetaDataServiceClient) error {

	mux.Handle("POST", pattern_MetaDataService_UpdateIptcData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/metadata_service.v1.MetaDataService/UpdateIptcData", runtime.WithHTTPPathPattern("/v1/iptc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MetaDataService_UpdateIptcData_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MetaDataService_UpdateIptcData_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MetaDataService_UpdateIptcData_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "iptc"}, ""))
)

var (
	forward_MetaDataService_UpdateIptcData_0 = runtime.ForwardResponseMessage
)
