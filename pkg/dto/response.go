package dto

import "github.com/google/uuid"

type Responsable interface {
	Comment | Country | User | Vote | NewUser | []*NewUser | SessionAuth |
		[]Comment | []Country | []User | []Vote | map[uuid.UUID]User | []Result
}

type Payload[T Responsable] struct {
	Body  T      `json:"body"`
	Error string `json:"error"`
}

func NewPayload[T Responsable](payload T, error string) Payload[T] {
	return Payload[T]{
		Body:  payload,
		Error: error,
	}
}
