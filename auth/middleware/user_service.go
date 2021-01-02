package middleware

import (
	"log"
	"users_api/db"
)

type UserService struct {
	store *db.UserStore
}

func CreateUserService(store *db.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (u *UserService) GetAllUsers() (users []db.User, err error) {
	users, err = u.store.GetAllUsers()

	if err != nil {
		log.Printf("Failed to fetch users: %s", err.Error())
		return nil, err
	}

	return users, nil
}

func (u *UserService) CreateUser(user *db.User) (err error) {
	err = u.store.CreateUser(user)

	if err != nil {
		log.Printf("Failed to create user: %s", err.Error())
		return
	}

	return
}
