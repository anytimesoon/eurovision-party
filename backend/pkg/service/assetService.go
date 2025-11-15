package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/anytimesoon/eurovision-party/pkg/errs"
)

type AssetService interface {
	PersistImage([]*multipart.FileHeader, string) *errs.AppError
}

type DefaultAssetService struct{}

func NewAssetService() DefaultAssetService {
	return DefaultAssetService{}
}

func (a DefaultAssetService) PersistImage(fileHeaders []*multipart.FileHeader, path string) *errs.AppError {
	file, err := fileHeaders[0].Open()
	if err != nil {
		log.Println("couldn't read file")
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fileName := fileHeaders[0].Filename

	f, err := os.Create(filepath.Join(path, fileName))
	if err != nil {
		log.Println("Failed to create image.", err)
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("Failed to create image.", err)
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}

	return nil
}
