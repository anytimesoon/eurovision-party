package router

import (
	"context"
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
		log.Printf("%s method %s was requested by %q", req.RequestURI, req.Method, req.RemoteAddr)
		next.ServeHTTP(resp, req)
	})
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		auth, appErr := authService.Authorize(token)
		if appErr != nil {
			log.Printf("%s method %s was requested by %q and rejected because token was rejected. %s", r.RequestURI, r.Method, r.RemoteAddr, appErr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "auth", *auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
