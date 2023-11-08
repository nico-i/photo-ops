package post

import (
	"github.com/nico-i/photo-ops/shared/go-shared/domain/image"
)

type Post struct {
	id       string
	image    image.Image
	caption  string
	hashtags []string
	location string
}
