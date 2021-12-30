package get_all_images

import (
	"strings"
	"testing"
	"uploader/entity"
	mock_entity "uploader/entity/mock"
	env "uploader/environment"
	"uploader/utils/formats"
	"uploader/utils/upload_status"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllImages(t *testing.T) {
	var expectedOutput []GetAllImagesDtoOutput
	var s3Url string

	if s3Url = env.S3_URL; strings.Compare(s3Url, "") == 0 {
		if s3Url = env.UPLOAD_URL; strings.Compare(s3Url, "") == 0 {
			s3Url = "http://localhost:3000/uploaded"
		}
	}
	img1 := GetAllImagesDtoOutput{
		ID:           "1",
		Name:         "Name",
		Size:         1.00,
		Extension:    formats.JPG,
		Status:       upload_status.UPLOADED,
		Url:          s3Url + "/1.jpg",
		ErrorMessage: "",
	}
	img2 := GetAllImagesDtoOutput{
		ID:           "2",
		Name:         "Name",
		Size:         2.00,
		Extension:    formats.JPG,
		Status:       upload_status.UPLOADED,
		Url:          s3Url + "/2.jpg",
		ErrorMessage: "",
	}
	expectedOutput = append(expectedOutput, img1, img2)

	var expectedReturn []entity.Image
	returnImg1 := entity.Image{
		ID:           "1",
		Name:         "Name",
		Size:         1.00,
		Extension:    formats.JPG,
		Status:       upload_status.UPLOADED,
		ErrorMessage: "",
	}
	returnImg2 := entity.Image{
		ID:           "2",
		Name:         "Name",
		Size:         2.00,
		Extension:    formats.JPG,
		Status:       upload_status.UPLOADED,
		ErrorMessage: "",
	}
	expectedReturn = append(expectedReturn, returnImg1, returnImg2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_entity.NewMockImageRepository(ctrl)
	repo.EXPECT().Get().Return(expectedReturn, nil)

	usecase := NewGetAllImages(repo)
	output, err := usecase.Execute()

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
