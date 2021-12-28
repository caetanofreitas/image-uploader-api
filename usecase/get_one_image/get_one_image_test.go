package get_one_image

import (
	"os"
	"strings"
	"testing"
	"uploader/entity"
	mock_entity "uploader/entity/mock"
	"uploader/utils/formats"
	"uploader/utils/upload_status"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetOneImage(t *testing.T) {
	input := "1"
	var s3Url string

	if s3Url = os.Getenv("S3_URL"); strings.Compare(s3Url, "") == 0 {
		if s3Url = os.Getenv("UPLOAD_URL"); strings.Compare(s3Url, "") == 0 {
			s3Url = "http://localhost:3000/uploaded"
		}
	}
	expectedOutput := GetOneImageDtoOutput{
		ID:           "1",
		Name:         "Name",
		Size:         1.00,
		Extension:    formats.JPG,
		Status:       upload_status.UPLOADED,
		Url:          s3Url + "/1.jpg",
		ErrorMessage: "",
	}

	expectedReturn := entity.NewImage()
	expectedReturn.ID = "1"
	expectedReturn.Name = "Name"
	expectedReturn.Size = 1.00
	expectedReturn.Extension = formats.JPG
	expectedReturn.Status = upload_status.UPLOADED
	expectedReturn.ErrorMessage = ""

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_entity.NewMockImageRepository(ctrl)
	repo.EXPECT().GetDetail(input).Return(*expectedReturn, nil)

	usecase := NewGetOneImage(repo)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
