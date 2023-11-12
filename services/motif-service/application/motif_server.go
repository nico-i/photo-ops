package application

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	gen "github.com/nico-i/photo-ops/services/motif-service/infrastructure/__generated__/motif_service"
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
	}

	_, b, _, _ := runtime.Caller(0)
	workDirPath := filepath.Dir(b)
	args := []string{fmt.Sprintf("%s/../../infrastructure/scripts/get_motif_bbox", workDirPath), imgPath}

	if debugDirPath != "" {
		args = append(args, "--debug", debugDirPath)
	}

	getBBoxScriptCmd := exec.Command("python", args...)

	var out bytes.Buffer
	getBBoxScriptCmd.Stdout = &out

	err := getBBoxScriptCmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error in running get_motif_bbox script: %w", err)
	}

	// parse out bytes as JSON
	res := gen.BoundingBoxResponse{}
	json.Unmarshal(out.Bytes(), &res)

	// remove tmp if it exists
	if _, err := os.Stat(tmpImgPath); !os.IsNotExist(err) {
		err = os.Remove(tmpImgPath)
		if err != nil {
			log.Fatal("error during the deletion of temp files: ", err)
		}
	}

	return &res, nil
}
