package router

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"net/http"
)

func writeResponse[T dto.Responsable](resp http.ResponseWriter, code int, data T, error string) {
	payload := dto.NewPayload(data, resp.Header().Get("Authorization"), error)
	resp.WriteHeader(code)
	if err := json.NewEncoder(resp).Encode(payload); err != nil {
		panic(err)
	}
}
