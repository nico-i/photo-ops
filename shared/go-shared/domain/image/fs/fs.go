package fs

import (
	"os"
	"path/filepath"

	"github.com/nico-i/photo-ops/shared/go-shared/domain/image"
)

type FileSystemRepository struct {
	root string
}

func New(root string) *FileSystemRepository {
	return &FileSystemRepository{root: root}
}

func (f *FileSystemRepository) Create(image image.Image) error {
	path := filepath.Join(f.root, image.GetName())
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(image.GetContent())
	if err != nil {
		return err
	}

	return nil
}

func (f *FileSystemRepository) Get(name string) (image.Image, error) {
	path := filepath.Join(f.root, name)
	file, err := os.Open(path)
	if err != nil {
		return image.Image{}, err
	}
	defer file.Close()

	content := make([]byte, 0)

	_, err = file.Read(content)
	if err != nil {
		return image.Image{}, err
	}

	return image.NewImage(name, content)
}

func (f *FileSystemRepository) Update(image image.Image) error {
	err := f.Delete(image.GetName())

	if err != nil {
		return err
	}

	return f.Create(image)
}

func (f *FileSystemRepository) Delete(name string) error {
	path := filepath.Join(f.root, name)

	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
