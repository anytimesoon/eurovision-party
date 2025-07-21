package data

import (
	"fmt"
	"github.com/anytimesoon/eurovision-party/conf"
	"github.com/anytimesoon/eurovision-party/pkg/data/dao"
	"github.com/timshannon/bolthold"
	"log"

	"github.com/google/uuid"
)

type VoteRepository interface {
	CreateVotes(uuid.UUID) error
	UpdateVote(dao.Vote) (*dao.Vote, error)
	GetVoteByUserAndCountry(uuid.UUID, string) (*dao.Vote, error)
	GetResults() (*[]dao.Result, error)
	GetResultsByUser(uuid.UUID) (*[]dao.Result, error)
	GetTotalVotesForCountry(string) (*dao.VoteTracker, error)
}

type VoteRepositoryDb struct {
	store *bolthold.Store
}

func NewVoteRepositoryDb(store *bolthold.Store) VoteRepositoryDb {
	return VoteRepositoryDb{store}
}

func (db VoteRepositoryDb) CreateVotes(userId uuid.UUID) error {
	countries := make([]dao.Country, 0)
	err := db.store.Find(&countries, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying country table during vote creation.", err)
		return err
	}

	for _, country := range countries {
		err = db.store.Upsert(
			voteKey(userId, country.Slug),
			dao.Vote{
				UserId:      userId,
				CountrySlug: country.Slug,
				Costume:     0,
				Song:        0,
				Performance: 0,
				Props:       0,
			},
		)
		if err != nil {
			log.Println("Error while inserting vote into vote table during vote creation.", err)
			return err
		}
	}

	return nil
}

func voteKey(userId uuid.UUID, countrySlug string) string {
	return fmt.Sprintf("%s_%s", userId.String(), countrySlug)
}

func (db VoteRepositoryDb) UpdateVote(vote dao.Vote) (*dao.Vote, error) {
	err := db.store.Update(
		voteKey(vote.UserId, vote.CountrySlug),
		vote,
	)
	if err != nil {
		log.Printf("Error when updating vote for user %s. %s", vote.UserId.String(), err)
		return nil, err
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetVoteByUserAndCountry(userId uuid.UUID, countrySlug string) (*dao.Vote, error) {
	var vote dao.Vote

	err := db.store.Get(voteKey(userId, countrySlug), &vote)
	if err != nil {
		log.Printf("Error when getting vote for user %s. %s", userId.String(), err)
		return nil, err
	}

	return &vote, nil
}

func (db VoteRepositoryDb) GetResults() (*[]dao.Result, error) {
	votes := make([]dao.Vote, 0)
	resultsMap := make(map[string]*dao.Result)

	err := db.store.Find(&votes, &bolthold.Query{})
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, err
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

func (db VoteRepositoryDb) GetResultsByUser(userId uuid.UUID) (*[]dao.Result, error) {
	votes := make([]dao.Vote, 0)
	results := make([]dao.Result, 0)

	err := db.store.Find(&votes, bolthold.Where("UserId").Eq(userId).Index("UserId"))
	if err != nil {
		log.Println("Error while querying vote table", err)
		return nil, err
	}

	for _, vote := range votes {
		results = append(results, vote.ToResult())
	}

	return &results, nil
}

func (db VoteRepositoryDb) GetTotalVotesForCountry(countrySlug string) (*dao.VoteTracker, error) {
	var voteCount dao.VoteTracker
	err := db.store.Get(countrySlug, &voteCount)
	if err != nil {
		log.Println("Error while querying vote count table", err)
		return nil, err
	}

	voteCount.Count++
	if voteCount.Count >= conf.App.VoteCountTrigger {
		voteCount.HasBeenNotified = true
	}

	err = db.store.Update(countrySlug, voteCount)
	if err != nil {
		log.Println("Error while updating vote count table", err)
		return nil, err
	}

	return &voteCount, nil
}
