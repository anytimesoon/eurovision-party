package router

import (
	"github.com/anytimesoon/eurovision-party/assets"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

type ImageHandler struct{}

func (ih ImageHandler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := filepath.Join(conf.App.Assets, params["file"])
	log.Println("serving file", fileName)
	http.ServeFile(w, r, fileName)
}

func assetHandler() http.Handler {
	fs := assets.NewStaticImageFS()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//f, err := subpath.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		//if err == nil {
		//	defer f.Close()
		//}
		//if os.IsNotExist(err) {
		//	r.URL.Path = "/"
		//}
		fs.ServeHTTP(w, r)
	})
}
