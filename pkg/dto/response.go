package dto

import "github.com/google/uuid"

type Responsable interface {
	Comment | *Country | *User | *Vote | *NewUser | []*NewUser |
		[]Comment | *[]Country | []User | []Vote | map[uuid.UUID]User
}

type Payload[T Responsable] struct {
	Session SessionAuth `json:"session"`
	Body    T           `json:"body"`
	Error   string      `json:"error"`
}

func NewPayload[T Responsable](payload T, auth Auth, error string) Payload[T] {
	return Payload[T]{
		Session: auth.ToSession(),
		Body:    payload,
		Error:   error,
	}
}
