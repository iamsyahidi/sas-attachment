package usecase

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)


type uploaderUseCase struct {
}

type UploaderUseCase interface {
	Upload(w http.ResponseWriter, r *http.Request) (bool, string)
	MultiUpload(w http.ResponseWriter, r *http.Request) (bool, []string)
}

func NewUploaderUseCase() UploaderUseCase {
	return &uploaderUseCase{}
}

// Upload 1 file
func (uploaderUseCase *uploaderUseCase) Upload(w http.ResponseWriter, r *http.Request) (bool, string)  {
	file, handler, err := r.FormFile("file")
	fileArray := strings.Split(handler.Filename, ".")
	fileName := fileArray[0] + "-" + time.Now().Format("20060102150405") + "." + fileArray[1]


	if err != nil {
		return false,fileName
	}

	defer file.Close()

	f, err := os.OpenFile("static/"+fileName,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("File "+fileName+" Fail uploaded")
		return false, fileName
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("File "+fileName+" Fail uploaded")
		return false, fileName
	}

	log.Println("File "+fileName+" Uploaded successfully")
	
	return true, fileName
}

func (uploaderUseCase *uploaderUseCase) MultiUpload(w http.ResponseWriter, r *http.Request) (bool, []string) {
	var error error
	var fileNames []string
	r.ParseMultipartForm(100000000)
	data := r.MultipartForm
	files := data.File["file"]
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			error = err
			break
		}

		defer file.Close()

		fileArray := strings.Split(files[i].Filename,".")
		fileName := fileArray[0]+"-"+time.Now().Format("20060102150405")+"."+fileArray[1]
		fileNames = append(fileNames, fileName)
		f, err := os.OpenFile("static/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("File "+fileName+" Fail uploaded")
			break
		}

		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			log.Println("File "+fileName+" Fail uploaded")
			break
		}

		log.Println("File "+fileName+" Uploaded successfully")
	}

	if error != nil {
		return false, fileNames
	}

	return true, fileNames
}