package cmd

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	// "os"
	"uploader/adapter/http/controllers/images_controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// func validateDb() {
// 	dbType := os.Getenv("DATABASE_TYPE")
// 	conn := os.Getenv("DATABASE_CONNECTION")
// 	db, err := sql.Open(dbType, conn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	_, err = db.Exec(`
// 		CREATE TABLE IF NOT EXISTS images
// 		(
// 			id TEXT NOT NULL,
// 			name TEXT NOT NULL,
// 			size REAL NOT NULL,
// 			extension TEXT NOT NULL,
// 			status TEXT NOT NULL,
// 			error_message TEXT,
// 			created_at    TEXT NOT NULL,
// 			updated_at    TEXT NOT NULL
// 		);
// 	`)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func configRoutes() {
	port := ":3000"
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Pong")
	}).Methods("GET")
	router.PathPrefix("/uploaded").Handler(http.StripPrefix("/uploaded", http.FileServer(http.Dir("upload"))))
	router.HandleFunc("/upload", images_controller.Upload).Methods("POST", http.MethodOptions)
	router.HandleFunc("/images", images_controller.GetAll).Methods("GET", http.MethodOptions)
	router.HandleFunc("/images/sse", images_controller.GetAllSSE).Methods("GET", http.MethodOptions)
	router.HandleFunc("/images/{id}", images_controller.GetDetail).Methods("GET", http.MethodOptions)
	fmt.Print("Sever launched 🚀 on port", port, "\n")
	log.Fatal(http.ListenAndServe(port, router))
}

func ExecuteRouter() {
	// validateDb()
	configRoutes()
}
