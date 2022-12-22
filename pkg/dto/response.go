package dto

type Responsable interface {
	Comment | *Country | *User | *Vote | *Auth | *EAuth |
		[]Comment | []Country | []User | []Vote
}

type Payload[T Responsable] struct {
	Token EAuth  `json:"token"`
	Body  T      `json:"body"`
	Error string `json:"error"`
}

func NewPayload[T Responsable](p T, token string, error string) Payload[T] {
	e := EAuth{
		EToken: token,
	}
	return Payload[T]{
		Token: e,
		Body:  p,
		Error: error,
	}
}
