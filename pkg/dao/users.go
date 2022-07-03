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
var userSlug string
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
		err = rows.Scan(&userID, &userName, &userSlug, &authLvl, &icon)
		if err != nil {
			log.Printf("FAILED to scan because %s", err)
			return users, err
		}
		users = append(users, domain.User{UUID: userID, Name: userName, Slug: userSlug, AuthLvl: authLvl, Icon: icon})
	}

	return users, nil
}

func User(user domain.User) (domain.User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`SELECT * FROM user WHERE name = '%s'`, user.Name)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return user, err
	}

	row := stmt.QueryRowContext(ctx)

	err = row.Scan(&userID, &userName, &userSlug, &authLvl, &icon)
	if err != nil {
		log.Println("scan FAILED!")
		return user, err
	}

	return domain.User{UUID: userID, Name: userName, Slug: userSlug, AuthLvl: authLvl, Icon: icon}, nil
}

func UsersUpdate(user domain.User, receivedUser domain.User) (domain.User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`UPDATE user SET name = '%s', icon = '%s' WHERE uuid = '%s'`, receivedUser.Name, receivedUser.Icon, receivedUser.UUID.String())
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return user, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! %s was not updated %s", user.Name, err)
		return user, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Printf("Error %s when finding rows affected", err)
		return user, err
	}

	log.Println("User rows affected:", rowsAffected)
	return receivedUser, nil
}
