package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// importing generated stubs
	gen "github.com/nico-i/photo-ops/services/motif-service/infrastructure/__generated__/motif_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MotifServerImpl struct {
	gen.UnimplementedMotifServiceServer
}

func (s *MotifServerImpl) GetBoundingBox(ctx context.Context, req *gen.ImageRequest) (*gen.BoundingBoxResponse, error) {
	fmt.Println("GetBoundingBox called")
	return &gen.BoundingBoxResponse{
		X:      1,
		Y:      2,
		Width:  3,
		Height: 4,
	}, nil
}

func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterMotifServiceServer(server, &MotifServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":9000"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
