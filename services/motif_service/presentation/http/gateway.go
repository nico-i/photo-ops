package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	gw "github.com/nico-i/photo-ops/services/motif-service/src/__generated__/motif_service"
)

var (
	serverAddr = flag.String("server", "localhost:9000", "gRPC server endpoint")
	port       = flag.String("p", "8080", "port to listen on")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMotifServiceHandlerFromEndpoint(ctx, mux, *serverAddr, opts)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("failed to register endpoint: %s", err))
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	portStr := fmt.Sprintf(":%s", *port)
	grpclog.Info("starting server on port %s", portStr)
	return http.ListenAndServe(portStr, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
