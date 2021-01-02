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

func (u *UserStore) GetAllUsers() (users []User, err error) {
	rows, err := u.db.Query("SELECT first_name, last_name, email FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users = []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
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
