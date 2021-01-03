package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"users_api/db"
	"users_api/middleware"

	"github.com/lib/pq"
)

type AuthApi struct {
	service *middleware.UserService
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func CreateAuthApi(service *middleware.UserService) *AuthApi {
	return &AuthApi{
		service: service,
	}
}

func (u *AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &db.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = u.service.CreateUser(user)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" { // Duplicate key error
				w.WriteHeader(http.StatusConflict)
				res, _ := json.Marshal(ErrorMessage{Message: err.Message})
				w.Write(res)
				return
			}

		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
