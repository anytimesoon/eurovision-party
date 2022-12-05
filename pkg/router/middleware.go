package router

import (
	"log"
	"net/http"
)

func jsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(resp, req)
	})
}

func imgHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "image/jpeg")
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(resp, req)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("%s was requested by %q", req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(resp, req)
	})
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header["Authorization"][0]
		newToken, err := authService.Authorize(h)
		if err != nil || newToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Add("Authorization", string(newToken))
		next.ServeHTTP(w, r)
	})
}
