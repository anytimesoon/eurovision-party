package data

import (
	"fmt"
	dto2 "github.com/anytimesoon/eurovision-party/pkg/api/dto"
	"github.com/anytimesoon/eurovision-party/pkg/api/enum"
	"github.com/anytimesoon/eurovision-party/pkg/errs"
	"github.com/anytimesoon/eurovision-party/pkg/service/dao"
	"github.com/timshannon/bolthold"
	"log"

	"github.com/google/uuid"
)

type VoteRepository interface {
	CreateVote(dto2.Vote) (*dao.Vote, *errs.AppError)
	UpdateVote(dto2.VoteSingle) (*dao.Vote, *errs.AppError)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dao.Vote, *errs.AppError)
	GetResults() (*[]dao.Result, *errs.AppError)
	GetResultsByUser(userId uuid.UUID) (*[]dao.Result, *errs.AppError)
}

type VoteRepositoryDb struct {
	store *bolthold.Store
}

func NewVoteRepositoryDb(store *bolthold.Store) VoteRepositoryDb {
	return VoteRepositoryDb{store}
}

func (db VoteRepositoryDb) CreateVote(voteDTO dto2.Vote) (*dao.Vote, *errs.AppError) {
	var vote dao.Vote

	vote = vote.FromDTO(voteDTO)
	err := db.store.Insert(
		voteKey(vote.UserId, vote.CountrySlug),
		vote,
	)
	if err != nil {
		log.Printf("Error when creating new vote for user %s. %s", voteDTO.UserId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotCreated + "your vote")
	}
	return &vote, nil
}

func voteKey(userId uuid.UUID, countrySlug string) string {
	return fmt.Sprintf("%s_%s", userId.String(), countrySlug)
}

func (db VoteRepositoryDb) UpdateVote(voteDTO dto2.VoteSingle) (*dao.Vote, *errs.AppError) {
	var vote dao.Vote

	err := db.store.Get(
		voteKey(voteDTO.UserId, voteDTO.CountrySlug),
		&vote)
	if err != nil {
		log.Printf("Error while fetching vote for user %s before update. %s", voteDTO.UserId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotUpdated + "your vote")
	}

	switch voteDTO.Cat {
	case enum.Song:
		vote.Song = voteDTO.Score
	case enum.Costume:
		vote.Costume = voteDTO.Score
	case enum.Performance:
		vote.Performance = voteDTO.Score
	case enum.Props:
		vote.Props = voteDTO.Score
	}

	vote.Total = int(vote.Costume) + int(vote.Song) + int(vote.Performance) + int(vote.Props)

	err = db.store.Update(
		voteKey(vote.UserId, vote.CountrySlug),
		vote,
	)

	return &vote, nil
}

func (db VoteRepositoryDb) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dao.Vote, *errs.AppError) {
	var vote dao.Vote

	err := db.store.Get(voteKey(userId, countrySlug), &vote)
	if err != nil && err.Error() == "No data found for this key" {
		log.Println("Found 0 votes from country and user. Creating a new vote")

		return db.CreateVote(dto2.Vote{
			UserId:      userId,
			CountrySlug: countrySlug,
		})
	} else if err != nil {
		log.Printf("Failed to find vote for country %s and user %s. %s", countrySlug, userId.String(), err)
		return nil, errs.NewUnexpectedError(errs.Common.NotFound + "your votes")
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetResults() (*[]dao.Result, *errs.AppError) {
	votes := make([]dao.Vote, 0)
	resultsMap := make(map[string]*dao.Result)

	err := db.store.Find(&votes, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, vote := range votes {
		if resultsMap[vote.CountrySlug] == nil {
			res := vote.ToResult()
			resultsMap[vote.CountrySlug] = &res
		} else {
			resultsMap[vote.CountrySlug].Costume += int(vote.Costume)
			resultsMap[vote.CountrySlug].Song += int(vote.Song)
			resultsMap[vote.CountrySlug].Performance += int(vote.Performance)
			resultsMap[vote.CountrySlug].Props += int(vote.Props)
			resultsMap[vote.CountrySlug].Total += int(vote.Costume) + int(vote.Song) + int(vote.Performance) + int(vote.Props)
		}
	}
	results := make([]dao.Result, 0, len(resultsMap))
	for _, res := range resultsMap {
		results = append(results, *res)
	}

	return &results, nil
}

func (db VoteRepositoryDb) GetResultsByUser(userId uuid.UUID) (*[]dao.Result, *errs.AppError) {
	votes := make([]dao.Vote, 0)
	results := make([]dao.Result, 0)

	err := db.store.Find(&votes, bolthold.Where("UserId").Eq(userId).Index("UserId"))
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, errs.NewUnexpectedError(errs.Common.DBFail)
	}

	for _, vote := range votes {
		results = append(results, vote.ToResult())
	}

	return &results, nil
}
