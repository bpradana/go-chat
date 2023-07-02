package common

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ParseForm(r *http.Request, formName string) (string, error) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		log.Println("[common] [ParseForm] error: ", err)
		return "", err
	}

	uploadedFile, handler, err := r.FormFile("profile_picture")
	if err != nil {
		log.Println("[common] [ParseForm] error: ", err)
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		log.Println("[common] [ParseForm] error: ", err)
		return "", err
	}

	profilePicture := handler.Filename
	fileLocation := filepath.Join(dir, "temp", profilePicture)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("[common] [ParseForm] error: ", err)
		return "", err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		log.Println("[common] [ParseForm] error: ", err)
		return "", err
	}

	return fileLocation, nil
}
