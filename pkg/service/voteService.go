package service

import (
	"encoding/json"
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"log"
)

type VoteService interface {
	CreateVote([]byte) (dto.Vote, error)
	UpdateVote([]byte) (dto.Vote, error)
}

type DefaultVoteService struct {
	repo domain.VoteRepository
}

func NewVoteService(repo domain.VoteRepository) DefaultVoteService {
	return DefaultVoteService{repo}
}

func (service DefaultVoteService) CreateVote(body []byte) (dto.Vote, error) {
	var voteDTO dto.Vote
	err := json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return voteDTO, err
	}

	vote, err := service.repo.CreateVote(voteDTO)
	if err != nil {
		log.Println("FAILED to create vote", err)
		return voteDTO, err
	}

	return vote.ToDto(), nil
}

func (service DefaultVoteService) UpdateVote(body []byte) (dto.Vote, error) {
	var voteDTO dto.Vote
	err := json.Unmarshal(body, &voteDTO)
	if err != nil {
		log.Println("FAILED to unmarshal json!", err)
		return voteDTO, err
	}

	vote, err := service.repo.UpdateVote(voteDTO)
	if err != nil {
		log.Println("FAILED to update vote", err)
		return voteDTO, err
	}

	return vote.ToDto(), nil
}
