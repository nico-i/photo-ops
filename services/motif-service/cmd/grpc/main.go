package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	// importing generated stubs
	gen "github.com/nico-i/photo-ops/services/motif-service/infrastructure/__generated__/motif_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	port = flag.String("port", "9000", "port to listen on")
)

const (
	tmpImgPath = "tmp.jpg"
)

type MotifServerImpl struct {
	gen.UnimplementedMotifServiceServer
}

func (s *MotifServerImpl) GetBoundingBox(ctx context.Context, req *gen.ImageRequest) (*gen.BoundingBoxResponse, error) {
	imgPath := req.GetPath()
	imgBase64 := req.GetBase64()
	debugDirPath := req.GetDebugDirPath()

	if imgBase64 != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			return nil, fmt.Errorf("error in decoding base64 string: %w", err)
		}

		if !strings.HasPrefix(string(decodedBytes), "\xff\xd8\xff") {
			return nil, fmt.Errorf("image is not of type jpg")
		}

		absoluteFilePath, err := filepath.Abs(tmpImgPath)
		if err != nil {
			return nil, fmt.Errorf("error in getting absolute file path: %w", err)
		}

		// Set the imgPath to the absolute file path of the temporary image
		imgPath = absoluteFilePath

		// Write the decoded bytes to the file
		err = os.WriteFile(imgPath, decodedBytes, 0666)
		if err != nil {
			return nil, fmt.Errorf("error in writing file: %w", err)
		}
	} else {
		imgPath = strings.ReplaceAll(imgPath, "\\", "/")

		// check if file exists
		if _, err := os.Stat(imgPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist")
		}
		// check if file is of type jpg
		if !strings.HasSuffix(imgPath, ".jpg") {
			return nil, fmt.Errorf("file is not of type jpg")
		}
	}

	args := []string{"../../infrastructure/scripts/get_bbox", imgPath}

	if debugDirPath != ""{
		args = append(args, "--debug", debugDirPath)
	}

	log.Println("executing python script with args: ", args)

	getBBoxScriptCmd := exec.Command("python", args...)

	var out bytes.Buffer
	getBBoxScriptCmd.Stdout = &out

	err := getBBoxScriptCmd.Run()
	if err != nil {
		log.Fatal("error during execution of python script: ", out.String(), err)
	}

	// parse out bytes as JSON
	res := gen.BoundingBoxResponse{}
	json.Unmarshal(out.Bytes(), &res)

	return &res, nil
}

func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gen.RegisterMotifServiceServer(server, &MotifServerImpl{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", fmt.Sprintf(":%s", *port)); err != nil {
		log.Fatal(fmt.Sprintf("error in listening on port :%s", *port), err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
