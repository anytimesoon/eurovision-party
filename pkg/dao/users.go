package dao

import (
	"context"
	db "eurovision/db"
	domain "eurovision/pkg/domain"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

var userID uuid.UUID
var authLvl domain.AuthLvl
var userName string
var icon string

func Users() ([]domain.User, error) {
	var users []domain.User
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	stmt, err := db.Conn.PrepareContext(ctx, "SELECT * FROM user")
	if err != nil {
		fmt.Println("FAILED to build query!")
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("FAILED to build query because %s", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userID, &userName, &authLvl, &icon)
		if err != nil {
			log.Printf("FAILED to scan because %s", err)
			return users, err
		}
		users = append(users, domain.User{UUID: userID, Name: userName, AuthLvl: authLvl, Icon: icon})
	}

	return users, nil
}
