package db

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AccountsStore struct {
	db *sql.DB
}

func CreateAccountsStore(conn *sql.DB) *AccountsStore {
	return &AccountsStore{
		db: conn,
	}
}

func (u *AccountsStore) CreateUser(email string, password string, firstName string, lastName string) (pwdHash string, err error) {
	pwdHash, err = u.hashPassword(password)

	if err != nil {
		log.Printf("Failed to create new user: %s", err.Error())
		return "", err
	}

	_, err = u.db.Exec(
		"INSERT INTO users(email, password, first_name, last_name) VALUES($1, $2, $3, $4)",
		email,
		pwdHash,
		firstName,
		lastName,
	)

	return pwdHash, err
}

func (u *AccountsStore) hashPassword(password string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pwdHash), err
}
