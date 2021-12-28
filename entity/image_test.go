package entity

import (
	"testing"
	"uploader/utils/errors_messages"

	"github.com/stretchr/testify/assert"
)

func TestImageWithSizeGreatherThan10MB(t *testing.T) {
	image := NewImage()
	image.ID = "1"
	image.Name = "New Image"
	image.Size = 1e8
	image.Extension = "jpg"

	err := image.IsValid()

	assert.Error(t, err)

	assert.Equal(t, errors_messages.SIZE_GREATHER_THAN_10MB, err.Error())
}

func TestImageWithInvalidExtension(t *testing.T) {
	image := NewImage()
	image.ID = "1"
	image.Name = "New Image"
	image.Size = 0.01
	image.Extension = "pdf"

	err := image.IsValid()

	assert.Error(t, err)

	assert.Equal(t, errors_messages.INVALID_IMAGE_EXTENSION, err.Error())
}

func TestValidImage(t *testing.T) {
	image := NewImage()
	image.ID = "1"
	image.Name = "New Image"
	image.Size = 0.01
	image.Extension = "jpeg"

	err := image.IsValid()
	assert.Nil(t, err)
}
