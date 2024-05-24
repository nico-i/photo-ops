package main

import (
	"flag"
	"fmt"
	"net"

	app "github.com/nico-i/photo-ops/services/motif_service/application/bbox"
	gen "github.com/nico-i/photo-ops/services/motif_service/application/bbox"

	// importing generated stubs
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	port = flag.String("p", "9000", "port to listen on")
)

func main() {
	flag.Parse()
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterBBoxServiceServer(server, &app.BBoxServiceImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	// start listening on port :8080 for a tcp connection
	portStr := fmt.Sprintf(":%s", *port)
	if l, err := net.Listen("tcp", portStr); err != nil {
		grpclog.Fatal(fmt.Sprintf("error in listening on port %s", portStr), err)
	} else {
		// the gRPC server
		grpclog.Info("starting server on port %s", portStr)
		if err := server.Serve(l); err != nil {
			grpclog.Fatal("unable to start server", err)
		}
	}
}
