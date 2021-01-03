package db

import (
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
}

func CreateUserStore(conn *sql.DB) *UserStore {
	return &UserStore{
		db: conn,
	}
}

func (u *UserStore) CreateUser(user *User) (err error) {
	_, err = u.db.Exec(
		"INSERT INTO users(email, password, first_name, last_name) VALUES($1, $2, $3, $4)",
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
	)
	return
}
