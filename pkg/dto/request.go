package dto

type Requestable interface {
	Comment | Country | User | Vote | Auth | EAuth | string
}

type Req[T Requestable] struct {
	Token string `json:"token"`
	Body  T      `json:"body"`
}
