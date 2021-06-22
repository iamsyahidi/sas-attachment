package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sas-attachment/model"
	"sas-attachment/usecase"
)

// uploaderHandler ...
type uploaderHandler struct {
	uploaderUseCase usecase.UploaderUseCase
}

// UploaderHandler ...
type UploaderHandler interface {
	Upload(w http.ResponseWriter, r *http.Request)
	MultiUpload(w http.ResponseWriter, r *http.Request)
	FileHandler(handl http.Handler) http.Handler
}

// NewUploaderHandler ...
func NewUploaderHandler(uploaderUseCase usecase.UploaderUseCase) UploaderHandler {
	return &uploaderHandler{
		uploaderUseCase: uploaderUseCase,
	}
}

func (handler *uploaderHandler) Upload(w http.ResponseWriter, r *http.Request) {
	success, filename := handler.uploaderUseCase.Upload(w, r)
	result := model.Result{
		IsSuccess: success,
		FileName:  filename,
	}
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("ERROR ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (handler *uploaderHandler) MultiUpload(w http.ResponseWriter, r *http.Request) {
	success, filenames := handler.uploaderUseCase.MultiUpload(w, r)
	result := model.Results{
		IsSuccess: success,
		FileNames: filenames,
	}
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("ERROR ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (handler *uploaderHandler) FileHandler(handl http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL))
		handl.ServeHTTP(w, r)
	})
}