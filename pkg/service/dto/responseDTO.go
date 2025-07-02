package dto

import "github.com/google/uuid"

type ApiResponseBody interface {
	Comment | Country | *User | Vote | *NewUser | []*NewUser | *Session |
		[]Comment | []Country | []User | []Vote | map[uuid.UUID]User | []Result
}

type ApiPayload[T ApiResponseBody] struct {
	Body  T      `json:"body"`
	Error string `json:"error"`
}

func NewApiPayload[T ApiResponseBody](payload T, error string) ApiPayload[T] {
	return ApiPayload[T]{
		Body:  payload,
		Error: error,
	}
}
