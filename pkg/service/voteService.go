package service

import (
	"eurovision/pkg/domain"
	"eurovision/pkg/dto"
	"eurovision/pkg/errs"
	"github.com/google/uuid"
)

//go:generate mockgen -source=voteService.go -destination=../../mocks/service/mockVoteService.go -package=service eurovision/pkg/service
type VoteService interface {
	//CreateVote(dto.Vote) (*dto.Vote, *errs.AppError)
	UpdateVote(dto.VoteSingle) (*dto.Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dto.Vote, *errs.AppError)
}

type DefaultVoteService struct {
	repo domain.VoteRepository
}

func NewVoteService(repo domain.VoteRepository) DefaultVoteService {
	return DefaultVoteService{repo}
}

//func (service DefaultVoteService) CreateVote(voteDTO dto.Vote) (*dto.Vote, *errs.AppError) {
//	appErr := voteDTO.Validate()
//	if appErr != nil {
//		return nil, appErr
//	}
//
//	vote, appErr := service.repo.CreateVote(voteDTO)
//	if appErr != nil {
//		return nil, appErr
//	}
//
//	result := vote.ToDto()
//	return &result, nil
//}

func (service DefaultVoteService) UpdateVote(voteSingleDTO dto.VoteSingle) (*dto.Vote, *errs.AppError) {
	appErr := voteSingleDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, appErr := service.repo.UpdateVote(voteSingleDTO)
	if appErr != nil {
		return nil, appErr
	}

	result := vote.ToDto()
	return &result, nil
}

func (service DefaultVoteService) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dto.Vote, *errs.AppError) {
	vote, _ := service.repo.GetVoteByUserAndCountry(userId, countrySlug)

	result := vote.ToDto()
	return &result, nil
}
