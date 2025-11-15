package api

import (
	"encoding/json"
	"net/http"

	"github.com/anytimesoon/eurovision-party/pkg/service/dto"
)

func WriteResponse[T dto.ApiResponseBody](resp http.ResponseWriter, code int, data T, error string) {
	payload := dto.NewApiPayload(data, error)

	resp.WriteHeader(code)
	err := json.NewEncoder(resp).Encode(payload)
	if err != nil {
		panic(err)
	}
}
