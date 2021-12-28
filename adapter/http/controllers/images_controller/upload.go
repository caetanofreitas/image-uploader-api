package images_controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"uploader/adapter/repository"
	aws_uploader "uploader/adapter/uploader/aws"
	local_uploader "uploader/adapter/uploader/local"
	"uploader/entity"
	"uploader/usecase/upload_image"
	"uploader/utils/errors_messages"

	uuid "github.com/nu7hatch/gouuid"
)

func getFile(r *http.Request) (upload_image.UploadImageDtoInput, error) {
	r.ParseMultipartForm(10 << 20) // 10MB
	file, fileHeader, err := r.FormFile("upload")
	if err != nil {
		return upload_image.UploadImageDtoInput{}, errors.New(errors_messages.UNABLE_TO_READ_FILE)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return upload_image.UploadImageDtoInput{}, errors.New(errors_messages.UNABLE_TO_READ_FILE)
	}

	Extension := strings.Split(fileHeader.Header["Content-Type"][0], "/")[1]

	ID, _ := uuid.NewV4()
	Name := fileHeader.Filename
	Size := float64(fileHeader.Size)

	input := upload_image.UploadImageDtoInput{
		ID:        ID.String(),
		Name:      Name,
		Size:      Size,
		Extension: Extension,
		Image:     fileBytes,
	}
	return input, nil
}

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "POST" {
		return
	}

	input, err := getFile(r)
	if err != nil {
		ErrorCallBack(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := ConnectDb()
	repo := repository.NewImageRepository(db)

	useAws := os.Getenv("USE_AWS")

	var uploader entity.ImageUploader
	if v, err := strconv.ParseBool(useAws); v && err == nil {
		uploader = aws_uploader.NewAwsUploader(
			os.Getenv("AWS_REGION"),
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET"),
			os.Getenv("S3_NAME"),
		)
	} else {
		uploader = local_uploader.NewLocalUploader()
	}

	usecase := upload_image.NewUploadImage(repo, uploader, UploadChannel)
	output, err := usecase.Execute(input)

	if err != nil {
		ErrorCallBack(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
