package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	authApi *AuthApi
}

func InitRouter(authApi *AuthApi) *Router {
	return &Router{
		authApi: authApi,
	}
}

func (r *Router) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	router.HandleFunc("/signup", r.authApi.SignUp).Methods("POST")

	return router
}
