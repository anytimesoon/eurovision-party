package main

import (
	"context"
	"github.com/anytimesoon/eurovision-party/pkg/api"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"log"
	"net/http"
)

func JsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(resp, req)
	})
}

func ImgHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "image/jpeg")
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(resp, req)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("%s method %s was requested by %q", req.RequestURI, req.Method, req.RemoteAddr)
		next.ServeHTTP(resp, req)
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session")
		if err != nil {
			log.Println("No session cookie was found. Trying authorization header.")
			api.WriteResponse(w, r, http.StatusUnauthorized, &dto.User{}, errs.Common.Unauthorized)
			return
		}

		auth, appErr := authService.Authorize(session.Value)
		if appErr != nil {
			log.Printf("%s method %s was requested by %q and rejected because token was rejected. %s", r.RequestURI, r.Method, r.RemoteAddr, appErr)
			api.WriteResponse(w, r, http.StatusUnauthorized, &dto.User{}, errs.Common.Unauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "auth", *auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
