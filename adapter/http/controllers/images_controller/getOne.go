package images_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"uploader/adapter/repository"
	"uploader/usecase/get_one_image"
	"uploader/utils/errors_messages"

	"github.com/gorilla/mux"
)

func getIdParam(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if strings.Compare(id, " ") == 0 {
		return id, errors.New(errors_messages.INVALID_IMAGE_ID)
	}

	return id, nil
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := getIdParam(r)
	if err != nil {
		ErrorCallBack(w, err.Error(), http.StatusNotFound)
		return
	}

	db := ConnectDb()
	defer db.Close()
	repo := repository.NewImageRepository(db)
	usecase := get_one_image.NewGetOneImage(repo)
	output, err := usecase.Execute(id)

	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == errors_messages.IMAGE_NOT_FOUND {
			status = http.StatusNotFound
		}

		ErrorCallBack(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
