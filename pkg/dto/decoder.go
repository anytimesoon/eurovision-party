package dto

import (
	"encoding/json"
	"log"
)

type Decodable interface {
	Country | User | Vote | UserImage | VoteSingle
}

func Decode[T Decodable](body []byte) (*T, error) {
	var model T
	log.Printf("%s", body)
	err := json.Unmarshal(body, &model)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return nil, err
	}
	return &model, nil
}
