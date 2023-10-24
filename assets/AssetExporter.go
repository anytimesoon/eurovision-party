package assets

import (
	"embed"
	"net/http"
)

var (
	//go:embed all:img
	img embed.FS

	//go:generate npm i
	//go:generate npm run build
	//go:embed all:svelteBuild
	SvelteUI embed.FS
)

func NewStaticImageFS() http.Handler {
	return http.FileServer(http.FS(img))
}

func NewSvelteUIFS() http.Handler {
	return http.FileServer(http.FS(SvelteUI))
}
