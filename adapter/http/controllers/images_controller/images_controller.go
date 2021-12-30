package images_controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	env "uploader/environment"
	"uploader/usecase/upload_image"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var UploadChannel = make(chan upload_image.UploadImageDtoOutput)

func ConnectDb() *sql.DB {
	dbType := env.DATABASE_TYPE
	conn := env.DATABASE_CONNECTION
	db, err := sql.Open(dbType, conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type errorResult struct {
	Message string `json:"error"`
}

func ErrorCallBack(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	output := errorResult{Message: errorMessage}
	json.NewEncoder(w).Encode(output)
}
