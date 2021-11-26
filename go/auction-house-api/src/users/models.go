package users

import (
	"auction-house/src/config"
	"auction-house/src/logger"
)

// User represents a person in application
type User struct {
	ID       string  `json:"id"`
	FullName string  `json:"fullName"`
	NickName string  `json:"nickName"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
}

//SaveUser create userin the database
func SaveUser(fullName string, nickName string, email string, balance float64) (err error) {
	if balance > 0 {
		_, err = config.DB.Exec("INSERT INTO users (full_name,nick_name,email,balance) VALUES ($1,$2,$3,$4)", fullName, nickName, email, balance)
	} else {
		_, err = config.DB.Exec("INSERT INTO users (full_name,nick_name,email) VALUES ($1,$2,$3)", fullName, nickName, email)
	}

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return err
	}
	return nil
}

// AllUsers retrieve all records os Item inside the data base
func AllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.FullName, &user.NickName, &user.Email, &user.Balance)
		if err != nil {
			logger.MongoLogger.Fatalf("server failed to handle: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return nil, err
	}
	return users, nil
}

// UserByID retrive a unique user by id
func UserByID(id string) (User, error) {
	row := config.DB.QueryRow("SELECT * FROM users WHERE user_id=$1", id)

	user := User{}
	err := row.Scan(&user.ID, &user.FullName, &user.NickName, &user.Email, &user.Balance)
	if err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return user, err
	}

	return user, nil
}
