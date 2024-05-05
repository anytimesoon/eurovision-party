package dto

import (
	"image"
	"mime/multipart"
)

type ProcessedImage struct {
	MultiPartImage *multipart.File
	File           image.Image
	FileExtension  string
	ID             string
}
