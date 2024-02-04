package assets

import (
	"embed"
	"net/http"
)

var (
	//go:embed img/newuser.png
	img embed.FS
)

func NewStaticAssetFS() http.Handler {
	return http.FileServer(http.FS(img))
}
