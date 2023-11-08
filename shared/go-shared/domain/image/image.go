package image

import (
	"errors"
	"strings"
)

var (
	InvalidImageName    = errors.New("Invalid image name")
	InvalidImageContent = errors.New("Invalid image content")
	InvalidImageType    = errors.New("Invalid image type")
)

type ImageType int64

const (
	JPEG ImageType = iota
	JPG
	PNG
)

type Image struct {
	name      string // unique
	content   []byte
	imageType ImageType
}

func (i Image) ImageTypeToString() string {
	switch i.imageType {
	case JPEG:
		return "jpeg"
	case JPG:
		return "jpg"
	case PNG:
		return "png"
	default:
		return ""
	}
}

func GetImageTypeByString(imageType string) (ImageType, error) {
	switch imageType {
	case "jpeg":
		return JPEG, nil
	case "jpg":
		return JPG, nil
	case "png":
		return PNG, nil
	default:
		return -1, InvalidImageType
	}
}

func (i Image) GetContent() []byte {
	return i.content
}

func (i Image) GetName() string {
	return i.name
}

func NewImage(name string, content []byte) (Image, error) {
	splitName := strings.Split(name, ".")

	if name == "" || len(splitName) != 2 {
		return Image{}, InvalidImageName
	}

	if content == nil || len(content) == 0 {
		return Image{}, InvalidImageContent
	}

	imgType, err := GetImageTypeByString(splitName[1])

	if err != nil {
		return Image{}, err
	}

	return Image{
		name:      name,
		content:   content,
		imageType: imgType,
	}, nil
}
