package router

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"log"
	"net/http"
)

func writeResponse[T dto.Responsable](resp http.ResponseWriter, req *http.Request, code int, data T, error string) {
	payload := dto.NewPayload(data, error)

	log.Printf("Sending %#v to %s", payload, req.RemoteAddr)
	resp.WriteHeader(code)
	err := json.NewEncoder(resp).Encode(payload)
	if err != nil {
		panic(err)
	}
}
