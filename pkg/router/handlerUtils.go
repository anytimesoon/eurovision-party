package router

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"log"
	"net/http"
)

func writeResponse[T dto.Responsable](resp http.ResponseWriter, req *http.Request, code int, data T, error string) {
	token := req.Context().Value("auth").(dto.Auth)
	payload := dto.NewPayload(data, token, error)
	log.Printf("Sending %#v to %s", payload, req.RemoteAddr)
	resp.WriteHeader(code)
	if err := json.NewEncoder(resp).Encode(payload); err != nil {
		panic(err)
	}
}
