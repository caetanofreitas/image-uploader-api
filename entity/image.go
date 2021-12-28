package entity

import (
	"errors"
	"uploader/utils/errors_messages"
	"uploader/utils/formats"
)

type Image struct {
	ID           string
	Name         string
	Size         float64
	Extension    string
	Status       string
	ErrorMessage string
}

func NewImage() *Image {
	return &Image{}
}

func ValidateImageSize(size float64) error {
	if size > 1e7 {
		message := errors_messages.SIZE_GREATHER_THAN_10MB
		return errors.New(message)
	}

	return nil
}

func ValidateImageExtension(extension string) error {
	err := formats.ValidateFormat(extension)

	if err != nil {
		message := errors_messages.INVALID_IMAGE_EXTENSION
		return errors.New(message)
	}

	return nil
}

func (i *Image) IsValid() error {
	err := ValidateImageSize(i.Size)

	if err != nil {
		return err
	}

	err = ValidateImageExtension(i.Extension)

	if err != nil {
		return err
	}

	return nil
}
