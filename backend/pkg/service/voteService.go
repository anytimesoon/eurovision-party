package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/anytimesoon/eurovision-party/conf"
	data2 "github.com/anytimesoon/eurovision-party/pkg/data"
	dao2 "github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/anytimesoon/eurovision-party/pkg/enum"
	"github.com/anytimesoon/eurovision-party/pkg/enum/chatMsgType"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/service/dto"
	"github.com/google/uuid"
)

type VoteService interface {
	UpdateVote(dto2.VoteSingle) (*dto2.Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dto2.Vote, *errs.AppError)
	GetResults() (*[]dto2.Result, *errs.AppError)
	GetResultsByUser(userId string) (*[]dto2.Result, *errs.AppError)
}

type DefaultVoteService struct {
	repo        data2.VoteRepository
	broadcast   chan dto2.SocketMessage
	commentRepo data2.CommentRepository
}

func NewVoteService(repo data2.VoteRepository, broadcast chan dto2.SocketMessage, commentRepo data2.CommentRepository) DefaultVoteService {
	return DefaultVoteService{
		repo,
		broadcast,
		commentRepo,
	}
}

func (vs DefaultVoteService) UpdateVote(voteSingleDTO dto2.VoteSingle) (*dto2.Vote, *errs.AppError) {
	appErr := voteSingleDTO.Validate()
	if appErr != nil {
		return nil, appErr
	}

	vote, err := vs.repo.GetVoteByUserAndCountry(voteSingleDTO.UserId, voteSingleDTO.CountrySlug)
	if err != nil {
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your votes")
	}

	switch voteSingleDTO.Cat {
	case enum.Song:
		vote.Song = voteSingleDTO.Score
	case enum.Costume:
		vote.Costume = voteSingleDTO.Score
	case enum.Performance:
		vote.Performance = voteSingleDTO.Score
	case enum.Props:
		vote.Props = voteSingleDTO.Score
	}

	vote.Total = int(vote.Costume) + int(vote.Song) + int(vote.Performance) + int(vote.Props)

	vote, err = vs.repo.UpdateVote(*vote)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	voteTracker, err := vs.repo.GetTotalVotesForCountry(voteSingleDTO.CountrySlug)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	if voteTracker.Count == conf.App.VoteCountTrigger && voteTracker.HasBeenNotified {
		go vs.broadcastVoting(voteTracker.ToDto())
	}

	result := vote.ToDto()
	return &result, nil
}

func (vs DefaultVoteService) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dto2.Vote, *errs.AppError) {
	vote, err := vs.repo.GetVoteByUserAndCountry(userId, countrySlug)
	if err != nil {
		return nil, errs.NewNotFoundError(errs.Common.NotFound + "your votes")
	}

	result := vote.ToDto()
	return &result, nil
}

func (vs DefaultVoteService) GetResults() (*[]dto2.Result, *errs.AppError) {
	results, err := vs.repo.GetResults()
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return sortResults(*results), nil
}

func (vs DefaultVoteService) GetResultsByUser(userId string) (*[]dto2.Result, *errs.AppError) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.BadlyFormedObject)
	}

	results, err := vs.repo.GetResultsByUser(id)
	if err != nil {
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	return sortResults(*results), nil
}

func sortResults(results []dao2.Result) *[]dto2.Result {
	sortedResultsDTO := make([]dto2.Result, 0)
	for _, result := range results {
		resultDTO := result.ToDto()

		inserted := false
		for i := range sortedResultsDTO {
			if resultDTO.Total > sortedResultsDTO[i].Total {
				sortedResultsDTO = append(sortedResultsDTO[:i+1], sortedResultsDTO[i:]...)
				sortedResultsDTO[i] = resultDTO
				inserted = true
				break
			}
		}

		if !inserted {
			sortedResultsDTO = append(sortedResultsDTO, resultDTO)
		}
	}

	return &sortedResultsDTO
}

func (vs DefaultVoteService) broadcastVoting(voteTracker dto2.VoteTracker) {
	comment := &dao2.Comment{
		UUID:      uuid.New(),
		UserId:    conf.App.BotId,
		Text:      fmt.Sprintf("People voted for %s %s", voteTracker.Country.Name, voteTracker.Country.Flag),
		FileName:  "",
		CreatedAt: time.Now(),
		ReplyTo:   nil,
	}

	comment, err := vs.commentRepo.CreateComment(*comment)
	if err != nil {
		log.Println("Failed to create comment for vote tracker.", err)
		return
	}

	commentDTO := comment.ToDto()
	commentDTO.IsVoteNotification = true
	voteTracker.Comment = commentDTO

	voteTrackerJson, err := json.Marshal(voteTracker)
	if err != nil {
		log.Println("Failed to marshal vote tracker.", err)
		return
	}

	log.Printf("Broadcasting vote tracker to all users for %s.", voteTracker.Country.Name)
	vs.broadcast <- dto2.SocketMessage{
		Category: chatMsgType.VOTE_NOTIFICATION,
		Body:     voteTrackerJson,
	}
}
