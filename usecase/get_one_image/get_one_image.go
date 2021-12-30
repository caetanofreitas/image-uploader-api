package get_one_image

import (
	"fmt"
	"strings"
	"uploader/entity"
	env "uploader/environment"
	"uploader/utils/upload_status"
)

type GetOneImage struct {
	Repository entity.ImageRepository
}

func NewGetOneImage(repo entity.ImageRepository) *GetOneImage {
	return &GetOneImage{Repository: repo}
}

func PrepareOutput(item entity.Image) GetOneImageDtoOutput {
	var s3Url string

	if s3Url = env.S3_URL; strings.Compare(s3Url, "") == 0 {
		if s3Url = env.UPLOAD_URL; strings.Compare(s3Url, "") == 0 {
			s3Url = "http://localhost:3000/uploaded"
		}
	}
	var Url string
	if item.Status != upload_status.UPLOADED {
		Url = ""
	} else {
		Url = fmt.Sprintf("%s/%s.%s", s3Url, item.ID, item.Extension)
	}

	image := GetOneImageDtoOutput{
		ID:           item.ID,
		Name:         item.Name,
		Size:         item.Size,
		Extension:    item.Extension,
		Status:       item.Status,
		Url:          Url,
		ErrorMessage: item.ErrorMessage,
	}

	return image
}

func (g *GetOneImage) Execute(id string) (GetOneImageDtoOutput, error) {
	result, err := g.Repository.GetDetail(id)

	if err != nil {
		return GetOneImageDtoOutput{}, err
	}

	output := PrepareOutput(result)
	return output, nil
}
