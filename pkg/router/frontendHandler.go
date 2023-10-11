package router

import (
	"eurovision/assets"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var svelteFS = assets.NewSvelteUIFS()

func frontendHandler() http.Handler {
	subpath, err := fs.Sub(assets.SvelteUI, "svelteBuild")
	if err != nil {
		log.Println("couldn't open the svelte subpath.", err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			svelteFS.ServeHTTP(w, r)
			return
		}
		f, err := subpath.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		svelteFS.ServeHTTP(w, r)
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
