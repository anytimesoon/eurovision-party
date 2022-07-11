package dao

import (
	"context"
	db "eurovision/db"
	"eurovision/pkg/dto"
	"eurovision/pkg/utils"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID    uuid.UUID `json:"id"`
	AuthLvl AuthLvl   `json:"authLvl"`
	Name    string    `json:"name"`
	Slug    string    `json:"slug"`
	Icon    string    `json:"icon"`
}

type AuthLvl int

const (
	None AuthLvl = iota
	Admin
)

func Users() ([]User, error) {
	var users []User
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
		var user User
		err = rows.Scan(&user.UUID, &user.Name, &user.Slug, &user.AuthLvl, &user.Icon)
		if err != nil {
			log.Printf("FAILED to scan because %s", err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func SingleUser(userSlug string) (User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var userDAO User

	query := fmt.Sprintf(`SELECT * FROM user WHERE slug = '%s'`, userSlug)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return userDAO, err
	}

	row := stmt.QueryRowContext(ctx)

	err = row.Scan(&userDAO.UUID, &userDAO.Name, &userDAO.Slug, &userDAO.AuthLvl, &userDAO.Icon)
	if err != nil {
		log.Println("scan FAILED!")
		return userDAO, err
	}

	return userDAO, nil
}

func CreateUser(userDTO dto.User) (User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var userDAO User

	userDTO.Data.Slug = utils.Slugify(userDTO.Data.Name)

	query := fmt.Sprintf(`INSERT INTO user(uuid, name, slug, authLvl, icon) VALUES ('%s', '%s', '%s', 0, '')`, uuid.New().String(), userDTO.Data.Name, userDTO.Data.Slug)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return userDAO, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! User was not created. %s", err)
		return userDAO, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return userDAO, err
	}
	log.Println("User rows affected:", rowsAffected)

	newUser, err := SingleUser(userDTO.Data.Slug)
	if err != nil {
		log.Printf("FAILED to find %s in database %s", userDTO.Data.Name, err)
		return userDAO, err
	}

	return newUser, nil
}

func UpdateUser(userDAO User, userDTO dto.User) (User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`UPDATE user SET name = '%s', icon = '%s' WHERE uuid = '%s'`, userDTO.Data.Name, userDTO.Data.Icon, userDAO.UUID.String())
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return userDAO, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! %s was not updated %s", userDAO.Name, err)
		return userDAO, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return userDAO, err
	}
	log.Println("User rows affected:", rowsAffected)

	newUser, err := SingleUser(userDTO.Data.Slug)
	if err != nil {
		log.Printf("FAILED to find %s in database %s", userDTO.Data.Name, err)
		return userDAO, err
	}

	return newUser, nil
}

func DeleteUser(userDAO User) (string, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := fmt.Sprintf(`DELETE FROM user WHERE slug = '%s'`, userDAO.Slug)
	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return userDAO.Name, err
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Printf("sql execution FAILED! User was not created. %s", err)
		return userDAO.Name, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return userDAO.Name, err
	}
	log.Println("User rows affected:", rowsAffected)

	return userDAO.Name, nil
}
