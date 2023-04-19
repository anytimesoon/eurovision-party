package dto

import (
	"encoding/json"
	"log"
)

type Decodable interface {
	Country | User | Vote | UserImage
}

func Decode[T Decodable](body []byte) (*T, error) {
	var model T
	err := json.Unmarshal(body, &model)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, err
	}
	return &model, nil
}
