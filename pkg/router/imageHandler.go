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

type ImageHandler struct{}

var fs = assets.NewStaticImageFS()

func (ih ImageHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["file"] == "default" {
		serveDefaultAvatar(w, r)
		return
	}

	fileName := filepath.Join(conf.App.Assets, params["file"])

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
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
