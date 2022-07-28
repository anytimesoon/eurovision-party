package handler

import (
	"encoding/json"
	"net/http"
)

func writeResponse(resp http.ResponseWriter, code int, data interface{}) {
	resp.WriteHeader(code)
	if err := json.NewEncoder(resp).Encode(data); err != nil {
		panic(err)
	}
}
