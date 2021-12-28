package upload_image

import (
	"testing"
	mock_entity "uploader/entity/mock"
	"uploader/utils/formats"
	"uploader/utils/upload_status"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUploadImageWhenItIsValid(t *testing.T) {
	input := UploadImageDtoInput{
		ID:        "1",
		Name:      "Image Name",
		Size:      1.00,
		Extension: formats.JPG,
		Image:     []byte{},
	}
	inputChan := make(chan UploadImageDtoOutput)

	expectedOutput := UploadImageDtoOutput{
		ID:           "1",
		Name:         "Image Name",
		Status:       upload_status.PENDING,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_entity.NewMockImageRepository(ctrl)
	uploader := mock_entity.NewMockImageUploader(ctrl)
	repository.EXPECT().Insert(input.ID, input.Name, input.Size, input.Extension, upload_status.PENDING, "")

	usecase := NewUploadImage(repository, uploader, inputChan)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
