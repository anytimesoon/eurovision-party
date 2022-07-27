package domain

import (
	"eurovision/pkg/dto"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VoteRepositoryDb struct {
	client *sqlx.DB
}

func NewVoteRepositoryDb(db *sqlx.DB) VoteRepositoryDb {
	return VoteRepositoryDb{db}
}

func (db VoteRepositoryDb) CreateVote(voteDTO dto.Vote) (Vote, error) {
	var vote Vote

	uuid := uuid.New().String()

	query := fmt.Sprintf(`INSERT INTO vote(uuid, userId, countryId, costume, song, performance, props) VALUES ('%s', '%s', '%s', %d, %d, %d, %d)`, uuid, voteDTO.UserId, voteDTO.CountryId, voteDTO.Costume, voteDTO.Song, voteDTO.Performance, voteDTO.Props)

	_, err := db.client.NamedExec(query, vote)
	if err != nil {
		log.Println("Error when creating new vote:", err)
		return vote, err
	}

	query = fmt.Sprintf(`SELECT * FROM vote WHERE uuid = '%s'`, uuid)
	err = db.client.Get(&vote, query)
	if err != nil {
		log.Println("Error when fetching vote after create", err)
		return vote, err
	}

	return vote, nil
}
