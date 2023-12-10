package file

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

var (
	FilePath = "G://Quik/backend/upload/"
)

func SaveFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	syffix := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	filename := uuid.New().String() + "." + syffix
	link := FilePath + filename
	dst, err := os.Create(link)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return filename, nil
}

func IsExistFile(filename string) bool {
	if _, err := os.Stat(FilePath + filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
