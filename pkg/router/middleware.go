package router

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"io"
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
		var token string

		//if request is GET look for token in header, else look in request body
		if r.Method == http.MethodGet {
			token = r.Header.Get("Authorization")
		} else {
			var authDTO *dto.Auth
			body, err := io.ReadAll(r.Body)
			err = json.Unmarshal(body, authDTO)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			token = authDTO.Token
		}

		newToken, appErr := authService.Authorize(token)
		if appErr != nil || newToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Add("Authorization", string(newToken))
		next.ServeHTTP(w, r)
	})
}
