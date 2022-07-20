package dto

// import (
// 	"context"
// 	db "eurovision/db"
// 	"eurovision/pkg/dto"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/google/uuid"
// )

// type Vote struct {
// 	UUID        uuid.UUID `db:"id"`
// 	UserId      uuid.UUID `db:"userId"`
// 	CountryId   uuid.UUID `db:"countryId"`
// 	Costume     int8      `db:"costume"`
// 	Song        int8      `db:"song"`
// 	Performance int8      `db:"performance"`
// 	Props       int8      `db:"props"`
// }

// func SingleVote(uuid uuid.UUID) (Vote, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelFunc()

// 	var voteDAO Vote

// 	query := fmt.Sprintf(`SELECT * FROM vote WHERE uuid = '%s'`, uuid.String())
// 	stmt, err := db.Conn.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return voteDAO, err
// 	}

// 	row := stmt.QueryRowContext(ctx)

// 	err = row.Scan(&voteDAO.UUID, &voteDAO.UserId, &voteDAO.CountryId, &voteDAO.Costume, &voteDAO.Song, &voteDAO.Performance, &voteDAO.Props)
// 	if err != nil {
// 		log.Printf("scan FAILED! %s", err)
// 		return voteDAO, err
// 	}

// 	return voteDAO, nil
// }

// func CreateVote(voteDTO dto.Vote) (Vote, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelFunc()

// 	var voteDAO Vote
// 	newUUID := uuid.New()

// 	query := fmt.Sprintf(`INSERT INTO vote(uuid, userId, countryId, costume, song, performance, props) VALUES ('%s', '%s', '%s', %d, %d, %d, %d)`, newUUID.String(), voteDTO.Data.UserId, voteDTO.Data.CountryId, voteDTO.Data.Costume, voteDTO.Data.Song, voteDTO.Data.Performance, voteDTO.Data.Props)
// 	stmt, err := db.Conn.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return voteDAO, err
// 	}

// 	res, err := stmt.ExecContext(ctx)
// 	if err != nil {
// 		log.Printf("sql execution FAILED! vote was not created. %s", err)
// 		return voteDAO, err
// 	}

// 	rowsAffected, err := res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return voteDAO, err
// 	}
// 	log.Println("Vote rows affected:", rowsAffected)

// 	newVote, err := SingleVote(newUUID)
// 	if err != nil {
// 		log.Printf("FAILED to find vote %s in database %s", newUUID, err)
// 		return voteDAO, err
// 	}

// 	return newVote, nil
// }

// func UpdateVote(voteDAO Vote, voteDTO dto.Vote) (Vote, error) {
// 	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelFunc()

// 	query := fmt.Sprintf(`UPDATE vote SET costume = %d, song = %d, performance = %d, props = %d WHERE uuid = '%s'`, voteDTO.Data.Costume, voteDTO.Data.Song, voteDTO.Data.Performance, voteDTO.Data.Props, voteDAO.UUID.String())
// 	stmt, err := db.Conn.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return voteDAO, err
// 	}

// 	res, err := stmt.ExecContext(ctx)
// 	if err != nil {
// 		log.Printf("sql execution FAILED! Vote was not updated %s", err)
// 		return voteDAO, err
// 	}

// 	rowsAffected, err := res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return voteDAO, err
// 	}
// 	log.Println("Vote rows affected:", rowsAffected)

// 	updatedVote, err := SingleVote(voteDTO.Data.UUID)
// 	if err != nil {
// 		log.Printf("FAILED to find vote in database %s", err)
// 		return voteDAO, err
// 	}

// 	return updatedVote, nil
// }
