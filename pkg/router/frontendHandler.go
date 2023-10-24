package router

import (
	"errors"
	"eurovision/assets"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var svelteFS = assets.NewSvelteUIFS()

func frontendHandler() http.Handler {
	fsys, err := fs.Sub(assets.SvelteUI, "svelteBuild")
	if err != nil {
		log.Println("couldn't open the svelte subpath.", err)
	}
	files := http.FS(fsys)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "svelteBuild")

		log.Printf("Serving page %s", path)
		//if path == "/" {
		//	svelteFS.ServeHTTP(w, r)
		//	return
		//}

		// try if file exists at path, if not append .html (SvelteKit adapter-static specific)
		_, err := files.Open(path)
		if errors.Is(err, os.ErrNotExist) {
			log.Println("adding html file extension")
			path = fmt.Sprintf("%s.html", path)
		}
		r.URL.Path = path
		http.FileServer(files).ServeHTTP(w, r)
	})
}

func frontend(w http.ResponseWriter, r *http.Request) {
	// If the requested file exists then return if; otherwise return index.html (fileserver default page)
	if r.URL.Path != "/" {
		fullPath := "svelteBuild" + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		_, err := os.Stat(fullPath)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
			// Requested file does not exist so we return the default (resolves to index.html)
			r.URL.Path = "/"
		}
	}
	svelteFS.ServeHTTP(w, r)
}
