package service

import (
	"bytes"
	"fmt"
	"github.com/anytimesoon/eurovision-party/pkg/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func cropImage(avatarDTO *dto.UserAvatar) (*dto.CroppedImage, *errs.AppError) {
	data, err := io.ReadAll(avatarDTO.File)
	buf := bytes.NewBuffer(data)
	img, fileExtension, err := image.Decode(buf)

	if err != nil {
		fmt.Println("Failed to decode image.", err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated)
	}

	return &dto.CroppedImage{
		File:          img.(subImager).SubImage(avatarDTO.CropBox),
		FileExtension: fileExtension,
		ID:            avatarDTO.UUID,
	}, nil
}

func storeImageToDisk(img *dto.CroppedImage) *errs.AppError {
	fileDir := filepath.Join(".", "assets", "img", "avatars")
	filePath := filepath.Join(fileDir, img.ID.String()+"."+img.FileExtension)

	err := os.MkdirAll(fileDir, 0750)
	if err != nil {
		log.Println("Couldn't create the avatar directory")
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create avatar file for user %s. %s", img.ID.String(), err)
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}

	switch img.FileExtension {
	case "jpg", "jpeg":
		err = jpeg.Encode(file, img.File, nil)
	case "png":
		err = png.Encode(file, img.File)
	default:
		log.Printf("Wrong file type when saving avatar for user %s. %s", img.ID.String(), err)
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}
	if err != nil {
		log.Printf("Failed to write avatar for user %s. %s", img.ID.String(), err)
		return errs.NewUnexpectedError(errs.Common.NotSaved)
	}

	return nil
}

//func stringToBin(base64Img string) (image.Image, *errs.AppError) {
//	imgString, err := base64.StdEncoding.DecodeString(base64Img)
//	if err != nil {
//		log.Println("Failed to decode image from base64.", err)
//		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "image")
//	}
//
//	reader := bytes.NewReader(imgString)
//
//	img, _, err := image.Decode(reader)
//	if err != nil {
//		log.Println("Failed to decode image from io reader.", err)
//		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "image")
//	}
//
//	return img, nil
//}
