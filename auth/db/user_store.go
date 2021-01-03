package db

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	db *sql.DB
}

func CreateUserStore(conn *sql.DB) *UserStore {
	return &UserStore{
		db: conn,
	}
}

func (u *UserStore) CreateUser(email string, password string, firstName string, lastName string) (string, string, error) {
	pwdHash, err := u.hashPassword(password)

	if err != nil {
		log.Printf("Failed to create new user: %s", err.Error())
		return "", "", err
	}

	_, err = u.db.Exec(
		"INSERT INTO users(email, password, first_name, last_name) VALUES($1, $2, $3, $4)",
		email,
		pwdHash,
		firstName,
		lastName,
	)

	return email, pwdHash, err
}

func (u *UserStore) hashPassword(password string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pwdHash), err
}
