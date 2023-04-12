package service

import (
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestCropImage(t *testing.T) {
	originalImageFile, err := os.Open(filepath.Join("..", "..", "assets", "img", "newuser.png"))
	if err != nil {
		panic(err)
	}
	defer originalImageFile.Close()

	originalImage, err := png.Decode(originalImageFile)
	if err != nil {
		panic(err)
	}

	CropImage("12345", originalImage)
}
