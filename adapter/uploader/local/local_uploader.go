package local_uploader

import (
	"bytes"
	"errors"
	"io"
	"os"
	comp "uploader/utils/image_compressor"
)

type LocalUploader struct{}

func NewLocalUploader() *LocalUploader {
	return &LocalUploader{}
}

func verifyAndCreateDir(id string, extension string) (string, error) {
	path := "upload"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	name := path + "/" + id + "." + extension
	return name, nil
}

func (l *LocalUploader) UploadImage(image []byte, id string, extension string) (int64, error) {
	compressedImage := comp.CompressImage(image)
	name, err := verifyAndCreateDir(id, extension)
	if err != nil {
		return 0, err
	}

	dst, err := os.Create(name)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, bytes.NewReader(compressedImage))
}
