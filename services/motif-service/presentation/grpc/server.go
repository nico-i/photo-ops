package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	app "github.com/nico-i/photo-ops/services/motif-service/application"
	gen "github.com/nico-i/photo-ops/services/motif-service/infrastructure/__generated__/motif_service"

	// importing generated stubs
	"google.golang.org/grpc"
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
	gen.RegisterMotifServiceServer(server, &app.MotifServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	// start listening on port :8080 for a tcp connection
	portStr := fmt.Sprintf(":%s", *port)
	if l, err := net.Listen("tcp", portStr); err != nil {
		log.Fatal(fmt.Sprintf("error in listening on port %s", portStr), err)
	} else {
		// the gRPC server
		log.Printf("starting server on port %s", portStr)
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
