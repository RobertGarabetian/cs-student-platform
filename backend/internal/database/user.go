package database

import (
	"cs-student-platform/backend/internal/models"
	"database/sql"
	"errors"
	"log"
)

func CreateUser(user models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = ?`
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		log.Println("Error fetching user:", err)
		return user, err
	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
