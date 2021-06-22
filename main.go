package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sas-attachment/handler"
	"sas-attachment/usecase"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	log.Println("WMS - Attachment Service starting...")

	r := mux.NewRouter()
	r.Use(CORSMiddleware)
	api := r.PathPrefix("/").Subrouter()

	uploaderUseCase := usecase.NewUploaderUseCase()
	uploaderHandler := handler.NewUploaderHandler(uploaderUseCase)

	api.HandleFunc("/upload", uploaderHandler.Upload).Methods("POST")
	api.HandleFunc("/multi-upload", uploaderHandler.MultiUpload).Methods("POST")
	r.PathPrefix("/").Handler(uploaderHandler.FileHandler(http.FileServer(http.Dir("static"))))

	log.Println("WMS - Attachment Service started!")

	port := ":" + os.Getenv("APP_PORT")
	fmt.Println(fmt.Sprintf("running on port %s !\n", port))
	server := http.Server{
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
		Handler:      r,
		Addr:         fmt.Sprintf("%s", port),
	}

	log.Fatal(server.ListenAndServe())
}

// CORSMiddleware ...
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//Enable CORS ...
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

		end := time.Now()
		executionTime := end.Sub(start)

		log.Println(
			r.RemoteAddr,
			r.Method,
			r.URL,
			r.Header.Get("user-agent"),
			executionTime.Seconds()*1000,
		)
	})
}
