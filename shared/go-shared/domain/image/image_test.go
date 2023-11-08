package image_test

import (
	"testing"

	"github.com/nico-i/photo-ops/shared/go-shared/domain/image"
)

func TestImage(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		content     []byte
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			content:     []byte("content"),
			expectedErr: image.InvalidImageName,
		},
		{
			test:        "No File Extension validation",
			name:        "name",
			content:     []byte("content"),
			expectedErr: image.InvalidImageName,
		},
		{
			test:        "Invalid File Extension validation",
			name:        "name.invalid",
			content:     []byte("content"),
			expectedErr: image.InvalidImageName,
		},
		{
			test:        "Nil Content validation",
			name:        "name",
			content:     nil,
			expectedErr: image.InvalidImageContent,
		},
		{
			test:        "Empty Content validation",
			name:        "name",
			content:     []byte{},
			expectedErr: image.InvalidImageContent,
		},
		{
			test:        "Valid Image",
			name:        "name.jpeg",
			content:     []byte("content"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := image.NewImage(tc.name, tc.content)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
