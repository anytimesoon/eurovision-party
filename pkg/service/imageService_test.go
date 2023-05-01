package service

import (
	"eurovision/pkg/dto"
	"github.com/google/uuid"
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

	newId := uuid.New()

	imageDTO := dto.UserImage{
		UUID:        newId,
		Image:       "",
		AuthLvl:     0,
		TopLeft:     40,
		TopRight:    600,
		BottomLeft:  40,
		BottomRight: 600,
	}

	path, appErr := cropImage(imageDTO, originalImage)
	if appErr != nil {
		t.Error("Expected no error, but got", appErr.Message)
	}

	expectedPath := filepath.Join("..", "..", "assets", "img", newId.String())
	if path != expectedPath {
		t.Errorf("Expected %s, but got %s", expectedPath, path)
	}
}
