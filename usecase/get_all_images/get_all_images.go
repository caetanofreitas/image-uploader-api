package get_all_images

import (
	"fmt"
	"os"
	"strings"
	"uploader/entity"
	"uploader/utils/upload_status"
)

type GetAllImages struct {
	Repository entity.ImageRepository
}

func NewGetAllImages(repo entity.ImageRepository) *GetAllImages {
	return &GetAllImages{Repository: repo}
}

func PrepareOutput(result []entity.Image) []GetAllImagesDtoOutput {
	var images []GetAllImagesDtoOutput
	var s3Url string

	if s3Url = os.Getenv("S3_URL"); strings.Compare(s3Url, "") == 0 {
		if s3Url = os.Getenv("UPLOAD_URL"); strings.Compare(s3Url, "") == 0 {
			s3Url = "http://localhost:3000/uploaded"
		}
	}

	for _, item := range result {
		var Url string
		if item.Status != upload_status.UPLOADED {
			Url = ""
		} else {
			Url = fmt.Sprintf("%s/%s.%s", s3Url, item.ID, item.Extension)
		}

		image := GetAllImagesDtoOutput{
			ID:           item.ID,
			Name:         item.Name,
			Size:         item.Size,
			Extension:    item.Extension,
			Status:       item.Status,
			Url:          Url,
			ErrorMessage: item.ErrorMessage,
		}
		images = append(images, image)
	}

	return images
}

func (g *GetAllImages) Execute() ([]GetAllImagesDtoOutput, error) {
	result, err := g.Repository.Get()

	if err != nil {
		return []GetAllImagesDtoOutput{}, err
	}

	output := PrepareOutput(result)
	return output, nil
}
