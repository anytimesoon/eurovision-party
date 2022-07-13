package dao

import (
	"context"
	db "eurovision/db"
	"eurovision/pkg/dto"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func Comments() ([]Comment, error) {
	var comments []Comment
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	stmt, err := db.Conn.PrepareContext(ctx, "SELECT * FROM comment")
	if err != nil {
		fmt.Println("FAILED to build query!")
		return comments, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println("rows FAILED!")
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.UUID, &comment.UserId, &comment.Text, &comment.CreatedAt)
		if err != nil {
			log.Println("scan FAILED!")
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func SingleComment(uuid uuid.UUID) (Comment, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var commentDAO Comment

	query := fmt.Sprintf(`SELECT * FROM comment WHERE uuid = '%s'`, uuid.String())
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return commentDAO, err
	}

	row := stmt.QueryRowContext(ctx)

	err = row.Scan(&commentDAO.UUID, &commentDAO.UserId, &commentDAO.Text, &commentDAO.CreatedAt)
	if err != nil {
		log.Printf("scan FAILED! %s", err)
		return commentDAO, err
	}

	return commentDAO, nil
}

func CreateComment(commentDTO dto.Comment) (Comment, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var commentDAO Comment
	newUUID := uuid.New()

	query := fmt.Sprintf(`INSERT INTO comment(uuid, userId, text) VALUES ('%s', '%s', '%s')`, newUUID.String(), commentDTO.Data.UserId, commentDTO.Data.Text)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return commentDAO, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! Comment was not created. %s", err)
		return commentDAO, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return commentDAO, err
	}
	log.Println("Comment rows affected:", rowsAffected)

	newComment, err := SingleComment(newUUID)
	if err != nil {
		log.Printf("FAILED to find comment %s in database %s", newUUID, err)
		return commentDAO, err
	}

	return newComment, nil
}
