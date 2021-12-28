package upload_image

import (
	"errors"
	"time"
	"uploader/entity"
	"uploader/utils/errors_messages"
	"uploader/utils/upload_status"
)

type UploadImage struct {
	Repository   entity.ImageRepository
	Uploader     entity.ImageUploader
	ResponseChan chan UploadImageDtoOutput
}

func NewUploadImage(repository entity.ImageRepository, uploader entity.ImageUploader, channel chan UploadImageDtoOutput) *UploadImage {
	return &UploadImage{
		Repository:   repository,
		Uploader:     uploader,
		ResponseChan: channel,
	}
}

func PrepareOutput(input UploadImageDtoInput, status string, error_message string) UploadImageDtoOutput {
	return UploadImageDtoOutput{
		ID:           input.ID,
		Name:         input.Name,
		Status:       status,
		ErrorMessage: error_message,
	}
}

func PrepareReturn(u *UploadImage, input UploadImageDtoInput, status string, error_message string) (UploadImageDtoOutput, error) {
	err := u.Repository.Insert(input.ID, input.Name, input.Size, input.Extension, status, error_message)

	if err != nil {
		return PrepareOutput(input, upload_status.FAILED, err.Error()), errors.New(errors_messages.DATABASE_ERROR)
	}

	output := PrepareOutput(input, status, error_message)
	return output, nil
}

func Upload(input UploadImageDtoInput, u *UploadImage) {
	if size, invalidImageUpload := u.Uploader.UploadImage(input.Image, input.ID, input.Extension); invalidImageUpload != nil {
		u.Repository.Update(input.ID, input.Name, input.Size, input.Extension, upload_status.FAILED, invalidImageUpload.Error())
		u.ResponseChan <- UploadImageDtoOutput{
			ID:           input.ID,
			Name:         input.Name,
			Status:       upload_status.FAILED,
			ErrorMessage: invalidImageUpload.Error(),
		}
	} else {
		u.Repository.Update(input.ID, input.Name, float64(size), input.Extension, upload_status.UPLOADED, "")
		u.ResponseChan <- UploadImageDtoOutput{
			ID:           input.ID,
			Name:         input.Name,
			Status:       upload_status.UPLOADED,
			ErrorMessage: "",
		}
	}
}

func (u *UploadImage) Execute(input UploadImageDtoInput) (UploadImageDtoOutput, error) {
	image := entity.NewImage()
	image.ID = input.ID
	image.Name = input.Name
	image.Size = input.Size
	image.Extension = input.Extension
	image.ErrorMessage = ""

	invalidInput := image.IsValid()

	if invalidInput != nil {
		output, err := PrepareReturn(u, input, upload_status.FAILED, invalidInput.Error())
		if err != nil {
			return output, err
		}

		return output, invalidInput
	}

	go func() {
		time.Sleep(5 * time.Second)
		go Upload(input, u)
	}()
	return PrepareReturn(u, input, upload_status.PENDING, "")
}
