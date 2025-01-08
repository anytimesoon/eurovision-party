package api

import (
	"errors"
	"github.com/anytimesoon/eurovision-party/assets"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AssetHandler struct {
	Service service.AssetService
}

var fs = assets.NewStaticAssetFS()

func (ah AssetHandler) CreateChatImage(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting chat image save")

	err := r.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		log.Println("Failed to parse form data", err)
		return
	}

	defer func(MultipartForm *multipart.Form) {
		err := MultipartForm.RemoveAll()
		if err != nil {

		}
	}(r.MultipartForm)

	fileHeaders := r.MultipartForm.File["file"]

	appErr := ah.Service.PersistImage(fileHeaders, filepath.Join(conf.App.Assets, "chat"))

	if appErr != nil {
		WriteResponse(w, appErr.Code, dto.Comment{}, appErr.Message)
	} else {
		WriteResponse(w, http.StatusOK, dto.Comment{}, "")
	}
}

func (ah AssetHandler) GetChatImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["file"] == "default" {
		log.Println("default avatar image requested")
		serveDefaultImage(w, r)
		return
	}

	fileName := filepath.Join(conf.App.Assets, "chat", params["file"])

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("couldn't find chat image. returning default")
		serveDefaultImage(w, r)
		return
	}
	log.Println("serving file", fileName)
	http.ServeFile(w, r, fileName)
}

func (ah AssetHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["file"] == "default" {
		log.Println("default avatar image requested")
		serveDefaultAvatar(w, r)
		return
	}

	fileName := filepath.Join(conf.App.Assets, "avatars", params["file"])

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("couldn't find avatar. returning default")
		serveDefaultAvatar(w, r)
		return
	}
	log.Println("serving file", fileName)
	http.ServeFile(w, r, fileName)
}

func DefaultAssetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "img/" + r.URL.Path
		fs.ServeHTTP(w, r)
	})
}

func serveDefaultAvatar(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "img/newuser.png"
	fs.ServeHTTP(w, r)
}

func serveDefaultImage(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "img/missingimage.png"
	fs.ServeHTTP(w, r)
}
