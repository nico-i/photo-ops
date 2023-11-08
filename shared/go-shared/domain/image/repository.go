package image

import "errors"

var (
	ErrImageNotFound = errors.New("Image not found")
	ErrCreateImage   = errors.New("Create image failed")
	ErrUpdateImage   = errors.New("Update image failed")
	ErrDeleteImage   = errors.New("Delete image failed")
)

type ImageRepository interface {
	Create(image Image) error
	Get(name string) (Image, error)
	Update(image Image) error
	Delete(name string) error
}
