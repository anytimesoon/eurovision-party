package router

import (
	"errors"
	"github.com/anytimesoon/eurovision-party/assets"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type AssetHandler struct{}

var fs = assets.NewStaticAssetFS()

func (ah AssetHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["file"] == "default" {
		log.Println("default avatar image requested")
		serveDefaultAvatar(w, r)
		return
	}

	fileName := filepath.Join(conf.App.Assets, params["file"])

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("couldn't find avatar. returning default")
		serveDefaultAvatar(w, r)
		return
	}
	log.Println("serving file", fileName)
	http.ServeFile(w, r, fileName)
}

func assetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func serveDefaultAvatar(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "img/newuser.png"
	fs.ServeHTTP(w, r)
}
