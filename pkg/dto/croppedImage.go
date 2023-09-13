package dto

import (
	"github.com/google/uuid"
	"image"
)

type CroppedImage struct {
	File          image.Image
	FileExtension string
	ID            uuid.UUID
}
