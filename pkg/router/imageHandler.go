package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

type ImageHandler struct{}

func (ih ImageHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := filepath.Join(".", "assets", "img", "avatars", params["file"])
	log.Println("serving file", fileName)
	http.ServeFile(w, r, fileName)
}
