package service

import (
	"bytes"
	"encoding/base64"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func cropImage(userImage dto.UserImage, originalImage image.Image) (string, *errs.AppError) {
	cropSize := image.Rect(userImage.TopLeft, userImage.BottomLeft,
		userImage.TopRight, userImage.BottomRight)
	croppedImage := originalImage.(subImager).SubImage(cropSize)

	path := filepath.Join("..", "..", "assets", "img", userImage.UUID.String())
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println("Failed to create new folder for user image.", err)
		return "", errs.NewUnexpectedError(errs.Common.NotCreated)
	}

	fileLocation := filepath.Join(path, "thumbnail.png")

	croppedImageFile, err := os.Create(fileLocation)
	if err != nil {
		log.Println("Failed to create new file for user image.", err)
		return "", errs.NewUnexpectedError(errs.Common.NotCreated + "image")
	}

	defer func(croppedImageFile *os.File) {
		err := croppedImageFile.Close()
		if err != nil {
			panic(err)
		}
	}(croppedImageFile)

	err = png.Encode(croppedImageFile, croppedImage)
	if err != nil {
		log.Println("Failed to encode new file for user image.", err)
		return "", errs.NewUnexpectedError(errs.Common.NotCreated + "image")
	}

	return fileLocation, nil
}

func stringToBin(base64Img string) (image.Image, *errs.AppError) {
	imgString, err := base64.StdEncoding.DecodeString(base64Img)
	if err != nil {
		log.Println("Failed to decode image from base64.", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "image")
	}

	reader := bytes.NewReader(imgString)

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Println("Failed to decode image from io reader.", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "image")
	}

	return img, nil
}
