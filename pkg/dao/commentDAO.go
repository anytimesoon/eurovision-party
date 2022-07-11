package dao

import (
	"context"
	db "eurovision/db"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID   uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
	Text   string    `json:"text"`
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
		err = rows.Scan(&comment.UUID, &comment.UserId, &comment.Text)
		if err != nil {
			log.Println("scan FAILED!")
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
