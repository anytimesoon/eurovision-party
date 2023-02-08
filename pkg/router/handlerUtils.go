package router

import (
	"encoding/json"
	"eurovision/pkg/dto"
	"net/http"
)

func writeResponse[T dto.Responsable](resp http.ResponseWriter, req *http.Request, code int, data T, error string) {
	token := req.Context().Value("authAndToken").(dto.AuthAndToken)
	payload := dto.NewPayload(data, token.Token, error)
	resp.WriteHeader(code)
	if err := json.NewEncoder(resp).Encode(payload); err != nil {
		panic(err)
	}
}
