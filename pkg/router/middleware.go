package router

import (
	"eurovision/pkg/domain"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type sessionStore struct {
	sessions map[string]session
}

type session struct {
	userId  uuid.UUID
	slug    string
	exp     time.Time
	authLvl domain.AuthLvl
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
		if token == nil {
			http.Error(w, "Please log in", http.StatusUnauthorized)
		} else {
			session, found := auth.sessions[token.Value]

			if found && session.exp.Before(time.Now()) {
				http.Error(w, "Your session has ended, please log in again", http.StatusUnauthorized)
			} else if found && session.exp.After(time.Now()) {
				session.exp = time.Now().Add(432000)
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Please log in", http.StatusUnauthorized)
			}
		}
	})
}
