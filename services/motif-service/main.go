package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	motifpb "github.com/nico-i/photo-ops/services/motif-service/infra/__generated__/go/motif"
)

type server struct {
	motifpb.UnimplementedMotifServiceServer
}

func NewServer() *server {
	return &server{}
}

func GetBoundingBox(context.Context, *motifpb.ImageRequest) (*motifpb.BoundingBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoundingBox not implemented")
}

func GetNoBGImage(context.Context, *motifpb.ImageRequest) (*motifpb.ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoBGImage not implemented")
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	motifpb.RegisterMotifServiceServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = motifpb.RegisterMotifServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
