package images_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"uploader/adapter/repository"
	"uploader/usecase/get_all_images"
	"uploader/utils/errors_messages"
)

type getAllRes struct {
	Images []get_all_images.GetAllImagesDtoOutput `json:"images"`
}

func formatOutput(output []get_all_images.GetAllImagesDtoOutput) *getAllRes {
	var images = []get_all_images.GetAllImagesDtoOutput{}
	if len(output) > 0 {
		images = output
	}

	return &getAllRes{
		Images: images,
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := ConnectDb()
	defer db.Close()

	repo := repository.NewImageRepository(db)
	usecase := get_all_images.NewGetAllImages(repo)
	output, err := usecase.Execute()

	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == errors_messages.IMAGE_NOT_FOUND {
			status = http.StatusNotFound
		}

		ErrorCallBack(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	formattedRes := formatOutput(output)
	json.NewEncoder(w).Encode(formattedRes)
}

func returnUpdateList(w http.ResponseWriter) {
	db := ConnectDb()
	defer db.Close()
	repo := repository.NewImageRepository(db)
	usecase := get_all_images.NewGetAllImages(repo)
	output, _ := usecase.Execute()

	formattedRes := formatOutput(output)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(formattedRes)
	fmt.Fprintf(w, "data: %v\n\n", buf.String())
}

func GetAllSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(1 * time.Second)
	select {
	case <-UploadChannel:
		returnUpdateList(w)
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
