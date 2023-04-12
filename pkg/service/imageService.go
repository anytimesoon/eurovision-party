package service

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func CropImage(userId string, originalImage image.Image) {
	bounds := originalImage.Bounds()
	width := bounds.Dx()
	//height := bounds.Dy()
	cropSize := image.Rect(0, 0, width/2+100, width/2+100)
	cropSize = cropSize.Add(image.Point{100, 80})
	croppedImage := originalImage.(SubImager).SubImage(cropSize)

	path := filepath.Join("..", "..", "assets", "img", userId)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	croppedImageFile, err := os.Create(filepath.Join(path, "thumbnail.png"))
	if err != nil {
		panic(err)
	}

	defer func(croppedImageFile *os.File) {
		err := croppedImageFile.Close()
		if err != nil {
			panic(err)
		}
	}(croppedImageFile)

	err = png.Encode(croppedImageFile, croppedImage)
	if err != nil {
		panic(err)
	}
}
