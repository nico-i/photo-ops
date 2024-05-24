package bbox

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vo "github.com/nico-i/photo-ops/services/motif_service/domain/value_objects"
)

const (
	tmpImgPath = "tmp.jpg"
)

type BBoxServiceImpl struct {
	UnimplementedBBoxServiceServer
}

func (s *BBoxServiceImpl) GetBBox(ctx context.Context, req *vo.Image) (*vo.BBox, error) {
	imgPath := req.GetPath()
	base64Str := req.GetBase64().Data

	if base64Str != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "error during the decoding of base64 string")
		}

		absoluteFilePath, err := filepath.Abs(tmpImgPath)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("error during the creation of absolute file path for temporary image: %s", err))
		}
		// Set the imgPath to the absolute file path of the temporary image
		imgPath = absoluteFilePath

		// Write the decoded bytes to the file
		err = os.WriteFile(imgPath, decodedBytes, 0666)
		if err != nil {
			return nil, status.Error(codes.Internal, "error during the creation of temporary image file for base64 string")
		}
	} else {
		imgPath = strings.ReplaceAll(imgPath, "\\", "/")
		// check if file exists
		if _, err := os.Stat(imgPath); os.IsNotExist(err) {
			return nil, status.Error(codes.InvalidArgument, "image does not exist")
		}
	}

	_, b, _, _ := runtime.Caller(0)
	workDirPath := filepath.Dir(b)
	pythonScriptPath := filepath.Join(workDirPath, "python/get_motif_bbox")
	// join workDirPath with python script path

	args := []string{pythonScriptPath, imgPath}
	if debugDirPath != "" {
		args = append(args, "--debug", debugDirPath)
	}

	getBBoxScriptCmd := exec.Command("python", args...)

	var out bytes.Buffer
	getBBoxScriptCmd.Stdout = &out

	err := getBBoxScriptCmd.Run()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error during the execution of python script: %s", err))
	}

	// parse out bytes as JSON
	res := gen.BoundingBoxResponse{}
	json.Unmarshal(out.Bytes(), &res)

	// remove tmp if it exists
	if _, err := os.Stat(tmpImgPath); !os.IsNotExist(err) {
		err = os.Remove(tmpImgPath)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("error during the deletion of temporary image file: %s", err))
		}
	}

	return &res, nil
}

func (s *BBoxServiceImpl) GetBBoxDebug(ctx context.Context, req *vo.Image) (*vo.BBox, error) {}
