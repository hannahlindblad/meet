package api

import (
	"auth/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lib/pq"
)

type AuthApi struct {
	store *db.UserStore
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type User struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func CreateAuthApi(store *db.UserStore) *AuthApi {
	return &AuthApi{
		store: store,
	}
}

func (u *AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(err.Error())
		return
	}

	user := &User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = u.store.CreateUser(
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
	)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" { // Duplicate key error
				w.WriteHeader(http.StatusConflict)
				res, _ := json.Marshal(ErrorMessage{Message: err.Message})
				w.Write(res)
				return
			}

		}
		log.Printf("Failed to create user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(User{FirstName: user.FirstName, LastName: user.LastName})
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
