package router

import (
	"log"
	"net/http"
)

type sessionStore struct {
	sessions map[string]string
}

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

func (auth sessionStore) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := r.Cookie("token")

		if _, found := auth.sessions[token.Value]; found {
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Please log in", http.StatusUnauthorized)
		}
	})
}
