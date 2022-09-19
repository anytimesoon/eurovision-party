package router

import (
	"log"
	"net/http"
	"time"
)

type sessionStore struct {
	sessions map[string]time.Time
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
		exp, found := auth.sessions[token.Value]
		log.Println("Expiration of token is at", exp)
		if found && exp.Before(time.Now()) {
			http.Error(w, "Your session has ended, please log in again", http.StatusUnauthorized)
		} else if found && exp.Before(time.Now()) {
			auth.sessions[token.Value] = time.Now().Add(432000)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Please log in", http.StatusUnauthorized)
		}
	})
}
