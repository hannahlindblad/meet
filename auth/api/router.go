package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	userApi *UserApi
}

func InitRouter(userApi *UserApi) *Router {
	return &Router{
		userApi: userApi,
	}
}

func (r *Router) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	router.HandleFunc("/users", r.userApi.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", r.userApi.CreateUser).Methods("POST")

	return router
}
