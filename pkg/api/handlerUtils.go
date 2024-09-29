package api

import (
	"encoding/json"
	"github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"net/http"
)

func WriteResponse[T dto.ApiResponseBody](resp http.ResponseWriter, req *http.Request, code int, data T, error string) {
	payload := dto.NewApiPayload(data, error)

	//log.Printf("Sending %#v to %s", payload, req.RemoteAddr)
	resp.WriteHeader(code)
	err := json.NewEncoder(resp).Encode(payload)
	if err != nil {
		panic(err)
	}
}
